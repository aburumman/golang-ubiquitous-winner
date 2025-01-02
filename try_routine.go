package main 

import (
	"fmt"
)

func greetings(s string) {
	fmt.Println("Hello There %s", s)
}

func main() {
	
	greetings("Patrick")
	go greetings("Moremi ")
}