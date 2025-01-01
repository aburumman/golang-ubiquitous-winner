package main 

import (
	"fmt"
	"os"
	"log"
)

func create_file(filename string) {

	new_file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error occured :", err)
		fmt.Errorf(err)
		log.Fatalf(err)
	}
}

func main() {

create_file()
}