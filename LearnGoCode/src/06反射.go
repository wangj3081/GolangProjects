package main

import (
	"fmt"
	"reflect"
)

func main06() {

	var a interface{} = 123
	var b interface{} = 456

	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	fmt.Printf("%T AND %T\n", t1, t2)
	if t1 == t2 {
		v1 := reflect.ValueOf(a).Int()
		v2 := reflect.ValueOf(b).Int()
		fmt.Println(v1 + v2)
	}

}
