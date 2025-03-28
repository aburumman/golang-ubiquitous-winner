package main 

import (
	"fmt"
)

func sum[T int64 | float64](x,y T)  T {
	return x * y
}

type MeanInput interface {
	int | uint | float32 | float64
}

func calcMean[T MeanInput](numbers []T)  (float64, error){
	var sum T
	for _, n := range numbers {
		sum += n
	} 
	return float64(sum), nil
}

func main() {
	fmt.Println(sum[int64](2,3))
}