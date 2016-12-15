package my_database

import (
	"reflect"
)

type Employee struct{
	Name string
	Surname string
	Position string
}

func (database *Employee) getCountOfFields() int{
	//home := new(Employee)
	//rcvr := reflect.ValueOf(home)
	//typ := reflect.Indirect(rcvr).Type()
	typ := reflect.TypeOf((*Employee)(nil)).Elem()
	return typ.NumField()
}