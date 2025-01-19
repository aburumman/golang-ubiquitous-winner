package main

import (
	"fmt"
)

func main() {

	var xy []int 
	var result string
	xy = []int{3, 26,37,48,59}

	for i := 0; i < len(xy); i++ {
		err := isOdd(xy[i])
		if err != nil {
			result += fmt.Sprintf("Error: %v\n", err) 
		} else {
			result += fmt.Sprintf("Number is odd\n")
		}
		


	}
fmt.Println("result: ", result)
}

func isOdd(x int) error{
	if x % 2 == 0 {
			return fmt.Errorf("Error: %d is even", x)
	}
	return nil
}