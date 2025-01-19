package main

import (
	"fmt"
	//"ioutil"
	"os"
)

var (
	i = 100
	f = 3.14
	b= true
	s = "Clear is better than clever"
	p = struct{x, y int64}{}
)

var bi int64 = 789
var ui uint64 = 34355

func main () {
	fmt.Printf("%v %v %v %v %v %v",  p.y , bi,  s, p.x, i, f,ui, b)
	myfile, err := OutPutTOWriter()
	if err != nil {
		fmt.Printf("There wasan error: %v", err)
		os.Exit(-1)
	}
	fmt.Printf("THe file name is: %s:", myfile)
}

func OutPutTOWriter() (string, error){

	file, err := os.CreateTemp("", "example")
	if err != nil {
		return "", err
	}
	defer file.Close()
	fmt.Fprintf(file, "%v %v %v %v %v %v",  p.y , bi,  s, p.x, i, f,ui, b )

	return file.Name(), nil
}