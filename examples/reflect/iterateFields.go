package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  uint8
}

func main() {
	p1 := person{"bob", 22}
	fmt.Printf("(%T) %+v\n", p1, p1)

	v := reflect.ValueOf(p1)
	t := v.Type()
	for i := range v.NumField() {
		value := v.Field(i)
		field := t.Field(i)

		fmt.Println(field.Name, field.Type, value)
	}

	// v := reflect.ValueOf(p1)
	// for i, field := range reflect.VisibleFields(reflect.TypeOf(p1)) {
	// 	fmt.Println(field.Name, field.Type, v.Field(i))
	// }
}
