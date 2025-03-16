package main

import (
	"os"
)

func main() {
	// _, err := os.Create("testfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	
	write_file := os.WriteFile("testfile.txt", []byte("This is the content"), 0644)
 if write_file != nil {
     panic(write_file)
 }
  //somebyte := []byte("This is the content")
 	read_file := os.ReadFile("testfile.txt", []byte(")

}