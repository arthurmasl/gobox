package main

import (
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Pos",
			Type: reflect.TypeOf(complex128(0)),
		},
		{
			Name: "Hp",
			Type: reflect.TypeOf(int(0)),
		},
	})

	v := reflect.New(typ)

	fmt.Printf("%#v\n", v)
	fmt.Printf("%#v\n", v.Elem().Field(0))
	fmt.Printf("%#v\n", v.Elem().Field(1))
}
