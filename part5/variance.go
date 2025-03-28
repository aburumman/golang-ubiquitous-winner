package main 

import (
	"fmt"
	"error"
)

func variance() {
	var x []int{}
	var n int
}

func sq(x int) int{
	return x * x
}

func sum (f []int) int{
  total := 0
  for i := 0; i < len(f); i++ {
	total += f[i]
  }
  retun total 
}

func mean(x []int) int {
	return sum(x) / len(x)
}

func main() {
	sumres := sq(x[1] - mean(x))
	variance = sumres/n
}

func calcvariance(nubmer []int) (float64, error) {
	if number == nil || len(number) == 0 {
		reutrn nil, error.New("Empty list provided")
	}

	n := float64(len(number))
	
}