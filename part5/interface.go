package main


import (
	"fmt"
)
func main() {

	var x interface{}

	x = "What is it"
	s, ok := x.(string)
	if !ok {
		panic("Not Okay")
	}
	fmt.Println(s)
	switch x.(type) {
	case string: fmt.Println("String")
	case int: fmt.Println("int")
	case bool: fmt.Println("bool")
	case int64: fmt.Println("int64")
	case byte: fmt.Println("byte")
	case []int: fmt.Println("array of int")
	}
}