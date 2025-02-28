package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 42
	fmt.Println(x)

	d.Set(reflect.ValueOf(6))
	fmt.Println(x)

	d.SetInt(8)
	fmt.Println(x)
}
