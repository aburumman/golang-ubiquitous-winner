package main

import (
	//"fmt"
	"log"
	"errors"
)

func isOdd(num int) error {
	if num % 2 == 0 {
		err := errors.New("Number is Odd")
		log.Println(err)
		return err
	}

	return nil //
}


func main(){

	for i := 1; i <= 10; i++ {
		isOdd(i)
	}
}