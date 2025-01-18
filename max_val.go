package main

import (
	"fmt"
)

func main() {
	var x int
	myList := []int {35, 34,54}
	if len(myList) == 1 {
		x = myList[0]
		fmt.Println( "Only one digit in the list ", x )
	} else if len(myList) > 1 {
		x = myList[0]
		for _,j := range myList {
		if j >= x {
			x = j
		}	
	}
}
fmt.Println("largest digit is %d", x)
}