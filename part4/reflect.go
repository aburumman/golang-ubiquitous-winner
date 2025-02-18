package main

import (
	"fmt"
	"reflect"
)

var (
	name = "Tomiwa"
	Age = 30
	height = 5.9
)

func main() {
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(Age))
	fmt.Println(reflect.TypeOf(height))
}