package main 

import (
	"fmt"
	//"os"
	//"flag"

)

func main() {
	zy := increment(6)
	xy := intCount(&zy)
	fmt.Println(xy.increase())
}

type increment int

type intCount interface {
	increase()	int
}

func (i *increment) increase() int {
	*i++
	return int(*i)
}