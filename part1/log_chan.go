package main

import (
	"fmt"
)


/****************************************************************

func main() {
	var misc string
	misc = <- isOdd(90)
 fmt.Println(misc)
}

func isOdd(n int) <- chan string {
	result := make(chan string)
	go func() {
	err := isEven(n)
	if err != nil {
		result <- fmt.Sprintf("Error: %v", err)
	} else {
		result <- fmt.Sprintf("%d is an odd number", n)
	}
	}()
	return result
}

func isEven(num int) error{

	if num % 2 == 0 {
		return fmt.Errorf("%d is an even number", num)
	}
	return nil
}





****************************************************************/

func main () {
	misc := <- isOdd(90)
	fmt.Println(misc)

}

func isEven(n int) error {
if n % 2 == 0 {
	return fmt.Errorf("Error: %d is even", n)
} 
	return nil
}

func isOdd(num int) <- chan string {

	result := make(chan string)

	go func() {
		if tally := isEven(num); tally != nil {
			result <- fmt.Sprintf("%v", tally)
		} else {
		result <- fmt.Sprintf("%d is Odd", num)
	}
		}()

	return result
}