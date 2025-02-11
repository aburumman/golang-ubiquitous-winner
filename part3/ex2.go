package main

import (
	"fmt"
	"sort"
)

func main() {
myList := []int{45, 67,12,34,79, 90}
sort.Ints(myList)
fmt.Println(myList)
}