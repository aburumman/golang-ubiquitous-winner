package main 

import (
	"fmt" //
)

func main() {

	constant (
		first := iota
		second int
		third int
		fourth int
	)
	fmt.Println(first, second, third, fourth)
}