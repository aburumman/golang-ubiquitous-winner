package main 

import (
	"fmt"
)

func smeArr(arr [7]int) []int {
	for x := 0; x < len(arr); x++ {
		arr[x] = x * x
	}
	return arr
}

func main() {
	myarr := smeArr([7]int{9,5,3,7,6,12})
	//myarr2 := smeArr([8]int{})
	fmt.Println(myarr)
}