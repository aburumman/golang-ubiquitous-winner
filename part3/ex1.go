package main  

import (
	"fmt"
)



func main() {
	arr1 := make([]int, 5)
	arr1 = append(arr1, 4,5,6,23,69)
	//arr2 := copy(arr1, arr2)
	//fmt.Println(arr2)
	fmt.Println(arr1)
	fmt.Println(add(34,78,90,2))
}

func add(value ...int) int {
	sum := 0
	for _, x := range value {
		sum += x
	} 
	return sum
}