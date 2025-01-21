package main

import (
	"fmt"
)

func main() {
  fmt.Println(isEven(24))
}

func isOdd(num int) {
panic(fmt.Sprintf("%d is not even", num))
}

func isEven(num int) bool {
	defer func () {
		if r := recover(); r != nil {
			fmt.Printf("Recovered %v\n", r)
		}
	}()

	if num % 2 != 0 {
isOdd(num)
	}
	return true
}