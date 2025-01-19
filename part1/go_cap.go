package main 

import (
	"fmt"
)

func myfunc(s []string) (string, int, int){

	gs := fmt.Sprintf("%#T", s)
	l := len(s)
	m := cap(s)

	return gs, m, l
}

func main() {
	fmt.Println(myfunc([]string{"hi", "you", "me"}))
}