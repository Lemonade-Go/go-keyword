package main

import (
	"fmt"
	"reflect"
)

func main() {

	//testing the implicit type when constant was assigned a number.
	//A numeric constant has no type until it’s given one, such as by an explicit cast.
	const n = 3.1415926
	x := 42
	y := float32(43.3)
	z := "hello"
	nt := reflect.TypeOf(n).Kind()
	xt := reflect.TypeOf(x).Kind()
	yt := reflect.TypeOf(y).Kind()
	zt := reflect.TypeOf(z).Kind()
	fmt.Println("%T: %s\n", nt, nt)
	fmt.Println("%T: %s\n", xt, xt)
	fmt.Println("%T: %s\n", yt, yt)
	fmt.Println("%T: %s\n", zt, zt)
	if nt == reflect.Float64 {
		println(">> n is float64")
	}
	if xt == reflect.Int {
		println(">> x is int")
	}
	if yt == reflect.Float32 {
		println(">> y is float32")
	}
	if zt == reflect.String {
		println(">> z is string")
	}
}
