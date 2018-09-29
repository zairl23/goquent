package goquent

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func (p *Product) TableName() string {
	return "products"
}