package main  

import (
	"fmt"
	"os"
	"errors"
)

func createFile(fileName string) (bool, error) {
	_, error := checkFile(fileName)
	if error != nil {
		return false, errors.New(error.Error())
	}
	fileCreated, err := os.Create(fileName)
	defer fileCreated.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}

func main() {
	fmt.Println(createFile("go_practice.go"))
}