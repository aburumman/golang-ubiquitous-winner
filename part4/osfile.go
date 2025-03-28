package main  

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("testfile.txt")

	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, 1024)
	output, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
	}
	//alist := []string(string(data[:output]))
	astring := string(data[:output])
	bstring := string(data[:100])
	fmt.Println(astring)
	fmt.Println(bstring)
	fmt.Println(output, string((data[:output])))
}