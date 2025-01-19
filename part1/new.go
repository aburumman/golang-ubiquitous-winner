package main

import (
	"fmt"
)
type Person struct {
	first, last string
}

func main(){
	people := make(map[Person]bool)
	people[Person{"John", "Cumming"}] = true
	people[Person{"Laide", "Ade"}] = false

	fmt.Println( people)
}