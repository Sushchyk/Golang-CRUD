package my_database

import (
	"reflect"
)

type Employee struct {
	Name     string
	Surname  string
	Position string
}

func (database *Employee) getCountOfFields() int {
	typ := reflect.TypeOf((*Employee)(nil)).Elem()
	return typ.NumField()
}