package main

import (
	"fmt"
)

func name_return()  (cvalid bool){
	var valid bool
	fmt.Println("%s", valid)
	valid = true
	fmt.Println("%s", valid)

	return 
}

func main() {

	b := name_return()
	fmt.Println("%s", b)
}