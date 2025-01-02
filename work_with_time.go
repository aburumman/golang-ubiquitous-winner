package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sub := time.Now().Sub(now)
	time.Sleep(20)
	since := time.Since(now)
	before := time.Now().Before(now)
	//now_now := time.Now
	later := time.Until(now)
	//later_since := time.Until(since)

	fmt.Println("Time now is: ", now)
	fmt.Println("Time sub is: ", sub)
	fmt.Println("Time since is: ", since)
	fmt.Println("Time before is: ", before)
	fmt.Println("Time Later is: ", later)
	//fmt.Println("Time Later is: ", later_since)
}