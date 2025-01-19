package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	now2 := time.Now().Local().UTC()
	second, x, y := time.Now().Date()
	sub := time.Now().Sub(now)
	time.Sleep(20)
	since := time.Since(now)
	before := time.Now().Before(now)
	//now_now := time.Now
	later := time.Until(now)
	//later_since := time.Until(since)

	fmt.Println("Time now is: ", now)
	fmt.Println("Time now2 is: ", now2)
	fmt.Println("Time second is: ", second, x, y)
	fmt.Println("Time sub is: ", sub)
	fmt.Println("Time since is: ", since)
	fmt.Println("Time before is: ", before)
	fmt.Println("Time Later is: ", later)
	//fmt.Println("Time Later is: ", later_since)
}