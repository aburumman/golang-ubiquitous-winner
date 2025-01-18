package main

import (
	"fmt"
	"math"
	"time"
	"math/rand"
	//"strconv"
	//"rand"
)

func main() {
	var some_num int 
	pl := fmt.Println 
	rand.Seed(time.Now().Unix())
	//some_num = strconv.ParseFloat(some_num)
	pl(math.Abs(float64(some_num)))
	pl(math.Sqrt(float64(some_num)))
	pl(math.Floor(float64(some_num)))
	pl(math.Log2(float64(some_num)))
	pl(math.Log2(float64(some_num)))
	pl(math.Pow(2,7))
	pl(math.Cbrt(float64(some_num)))
	pl(math.Max(6,-90))
	pl(math.Min(45,-8,))
}