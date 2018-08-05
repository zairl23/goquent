package goquent

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
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
		log.Fatalln("atabase connection failed")
		// fmt.Printf("Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
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
	return DB.Writer.Where(where).Assign(model).FirstOrCreate(model).Error
}