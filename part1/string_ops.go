package main

import (
	"fmt"
	"strings"
)

func main () {
	name := "Axel"

	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Contains(name, "A"))
	//var substitute string.NewReplacer
	substitute := strings.NewReplacer( "A", "P")
	name = substitute.Replace(name)
	fmt.Println(name)
	fmt.Println(strings.Index(name, "A"))
	my_sent := "This is the rule of the system"
	split_sent := strings.Split(my_sent, " ")
	fmt.Println(split_sent)
	fmt.Println(strings.HasPrefix(my_sent, "A"))
}