package goquent

import (
	"fmt"
	"testing"
	"github.com/zairl23/config"
    // _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestOne(t *testing.T) {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	//   panic("failed to connect database")
	// }
	// defer db.Close()
  
	// // Migrate the schema
	// db.AutoMigrate(&Product{})
  
	// Create
	// product := &Product{Code: "L1212", Price: 2000}
	// product.Create()

	// Find
	// p1 := Product{Code: "1"}
	// p1 := P()
	// fmt.Println(ProductModel.Database)
	// p1.init
	//p, err := p1.Find()

	// if err != nil {
	// 	panic(err)
	// }

	//fmt.Println(p1)
  
	// Read
	// var product Product
	// db.First(&product, 1) // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212
  
	// // Update - update product's price to 2000
	// //db.Model(&product).Update("Price", 2000)
  
	// fmt.Println(product)
	// Delete - delete product
  //   db.Delete(&product)
		 // init config
    if err := config.Init("config.yaml"); err != nil {
		panic(err)
	}

	DB.Init()
	
	// // Migrate the schema
	DB.Writer.AutoMigrate(&Product{})

	// Create
	product := &Product{Code: "L1212", Price: 2000}
	err := DB.Create(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product.ID, product.Price)

	// Update
	err = DB.Update(product, Product{Price: 2001})
	if err != nil {
		panic(err)
	}

	fmt.Print(product.Price)

	// ManyUpdate
	// err = DB.Update(Product{}, Product{Price: 2002})
	// if err != nil {
	// 	panic(err)
	// }

	// ManyUpdate with conditional query
	// err = DB.UpdateWhere(Product{}, "id = 1", Product{Price: 2003})
	// if err != nil {
	// 	panic(err)
	// }

	// Delete
	err = DB.Delete(product)
	if err != nil {
		panic(err)
	}

	// Many Delete with where conditional query
	err = DB.DeleteWhere(Product{}, "id < 3")
	if err != nil {
		panic(err)
	}

	// Query 
	a := make([]Product, 1)
	err = DB.Find(&a, "id = 9")
	if err != nil {
		panic(err)
	}

	fmt.Println(a)

	// Query one record
	b := Product{}
	err = DB.Find(&b, "id = 9")
	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	// UpdateOrCreate
	c := &Product{Price: 2005, Code: "L346"}
	err = DB.UpdateOrCreate(c, Product{Price: 2005})
	if err != nil {
		panic(err)
	}
	
	fmt.Println(*c)

}