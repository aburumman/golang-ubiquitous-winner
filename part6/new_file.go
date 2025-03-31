package main   

import (
	"fmt"
)

func main() {
	mystr := "Hello there"
	s := []byte(mystr)
	// for _, item := range s {
	// 	fmt.Println(string(item))
	// }
	//fmt.Println(len(s))
	//fmt.Printf("%v %T %t \n", s, s, s)
	fmt.Printf("%v %T %t %b\n", mystr, mystr, mystr, s)
	
}