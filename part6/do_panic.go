package main 

import (
	"fmt"
	"log"
	"errors"
)

func main() {
	fmt.Println("Starting...")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering...")
		} 
	}()
	panic("This code panicked")
	fmt.Println("Running...")
}

func panic_if_fail(err string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering...")
			log.Println(err)
		}
	}()
	panic(err)
}