package main 

import (
	"fmt"
)

func main() {
	var a int = 62
	var b *int = &a
	fmt.Println(*b, b)

	type structA struct {
		Name string
	}
	var some_inst *structA
	fmt.Println(some_inst)
	some_inst = new(structA)
	fmt.Println(some_inst)
	x := "Astron"
	y := "Blastron"
	sayGreeting(*x, *y)
	
	
}

func sayGreeting(first, last *string) {
	fmt.Println(*first, *last)

}