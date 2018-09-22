package goquent

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func (Product) TableName() string {
	return "products"
}