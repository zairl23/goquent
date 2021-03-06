package goquent

import (
	"fmt"
	"testing"
	"github.com/zairl23/config"
)

func TestOne(t *testing.T) {
	// init config
    if err := config.Init("config.yaml"); err != nil {
		panic(err)
	}

	DB.Init()
	
	// Migrate the schema
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
	err = DB.Update(Product{}, Product{Price: 2002})
	if err != nil {
		panic(err)
	}

	// ManyUpdate with conditional query
	err = DB.UpdateWhere(Product{}, "id = 1", Product{Price: 2003})
	if err != nil {
		panic(err)
	}

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

	// Count
	total := DB.Count(Product{}, "")

	fmt.Println(total)

	// Paginate
	products := make([]Product, 0)
	query := PaginateQuery{
		Model: Product{},
		Result: &products,
		Limit: 10,
		Page: 1,
		Where: "",
		Fields: "*",
		Order: "created_at desc",
	}

	totals, err := DB.Paginate(query)

	if err != nil {
		panic(err)
	}

	fmt.Println(products)

	fmt.Println(totals)

}