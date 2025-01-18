package main

import (
	"fmt"
)

func main() {
	var a, b int = 11,20
	yes := "Yes"
	fmt.Println(yes)

	if frac := float64(a)/float64(b); frac > 0.5 {
		fmt.Println("frac a/b is %v", frac)
	}

	x := 11.0
	y := 20.0

	if frac := float64(x) / float64(y); frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}