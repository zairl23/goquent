package goquent

import (

)

type Repository interface {
	Create(model interface{}) error
	Update(model interface{}, where string, data interface{}) error
	Take(model interface{}, where string) (interface{}, error)
	Delete(model interface{}, where string) error
	Count(model interface{}, where string) (uint64, error)
	All(PaginateQuery) (uint64, error)
}