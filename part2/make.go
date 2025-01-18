package main

import (
	"fmt"
)

func main() {
	anArray := make( []int, 5)
	secArray := make([]string, 2, 5)
	fmt.Println(anArray, secArray)
	c := b[:2]
	fmt.Println(c)
}