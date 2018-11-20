package goquent

import (
	"fmt"
	"log"

	// MySQL driver.
	"github.com/jinzhu/gorm"
	"github.com/zairl23/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Writer   *gorm.DB
	Reader   *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Fatalln("database connection failed")
		// fmt.Printf("Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(config.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000)
	db.DB().SetMaxIdleConns(0)
}

// used for cli
func InitWriterDB() *gorm.DB {
	return openDB(config.Get("db.username"),
		config.Get("db.password"),
		config.Get("db.addr"),
		config.Get("db.name"))
}

func GetWriterDB() *gorm.DB {
	return InitWriterDB()
}

func InitReaderDB() *gorm.DB {
	return openDB(config.Get("reader_db.username"),
		config.Get("reader_db.password"),
		config.Get("reader_db.addr"),
		config.Get("reader_db.name"))
}

func GetReaderDB() *gorm.DB {
	return InitReaderDB()
}

func (db *Database) Init() {
	DB = &Database{
		Writer:   GetWriterDB(),
		Reader: GetReaderDB(),
	}
}

func (db *Database) Close() {
	DB.Writer.Close()
	DB.Reader.Close()
}

func (d *Database) Create(model interface{}) error {
	return DB.Writer.Create(model).Error
}

func (d *Database) Update(model interface{}, data interface{}) error {
	return DB.Writer.Model(model).Updates(data).Error
}

func (d *Database) UpdateWhere(model interface{}, where string, data interface{}) error {
	return DB.Writer.Model(model).Where(where).Updates(data).Error
}

func (d *Database) Delete(model interface{}) error {
	return DB.Writer.Delete(model).Error
}

func (d *Database) DeleteWhere(model interface{}, where string) error {
	return DB.Writer.Where(where).Delete(model).Error
}

func (d *Database) Find(model interface{}, where string) error {
	return DB.Reader.Where(where).Find(model).Error
}

func (d *Database) First(model interface{}, where string) error {
	return DB.Reader.Where(where).First(model).Error
}

func (d *Database) UpdateOrCreate(model interface{}, where interface{}) error {
	err := DB.Writer.Where(where).First(where).Error

	if err != nil {
		err = DB.Writer.Create(model).Error
	} else {
		err = DB.Writer.Model(where).Updates(model).Error
	}

	return err
}

func (d *Database) Count(model interface{}, where string) uint64 {
	var total uint64
	var err error 
	
	if where != "" {
		err = DB.Reader.Model(model).Where(where).Count(&total).Error
	} else {
		err = DB.Reader.Model(model).Count(&total).Error
	}
	

	if err != nil {
		total = 0
	}

	return total
}

type PaginateQuery struct {
	Model interface{}
	Result interface{} // must be a address
	Limit uint64
	Offset uint64
	Where string
	Fields string
	Order string
	Group string
	Preloads map[string]interface{} // 预加载项
}

// result := make([]model, 0)
func (d *Database) Paginate(query PaginateQuery) (uint64, error) {
	var total uint64
	var err error
	total = 0

	if query.Order == "" {
		query.Order = "id desc"
	}

	if query.Where != "" {
		if len(query.Preloads) == 0 {
			err = DB.Reader.Model(query.Model).Where(query.Where).Select(query.Fields).Count(&total).Error
			err = DB.Reader.Model(query.Model).Where(query.Where).Select(query.Fields).Limit(query.Limit).Offset(query.Offset).Order(query.Order).Group(query.Group).Find(query.Result).Error
		} else {
			err = DB.Reader.Model(query.Model).Where(query.Where).Select(query.Fields).Count(&total).Error
			qu := DB.Reader.Model(query.Model)
			for key, fun := range query.Preloads {
				qu = qu.Preload(key, fun)
			}
			err = qu.Where(query.Where).Select(query.Fields).Limit(query.Limit).Offset(query.Offset).Order(query.Order).Group(query.Group).Find(query.Result).Error
		}
		
	} else {
		if len(query.Preloads) == 0 {
			err = DB.Reader.Model(query.Model).Select(query.Fields).Count(&total).Error
			err = DB.Reader.Model(query.Model).Select(query.Fields).Limit(query.Limit).Offset(query.Offset).Order(query.Order).Group(query.Group).Find(query.Result).Error
		} else {
			err = DB.Reader.Model(query.Model).Where(query.Where).Select(query.Fields).Count(&total).Error
			qu := DB.Reader.Model(query.Model)
			for key, fun := range query.Preloads {
				qu = qu.Preload(key, fun)
			}
			err = qu.Where(query.Where).Select(query.Fields).Limit(query.Limit).Offset(query.Offset).Order(query.Order).Group(query.Group).Find(query.Result).Error
		}
	}
	return total, err
}
