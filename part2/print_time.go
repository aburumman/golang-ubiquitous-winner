package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	pr := fmt.Println
	hour := now.Hour()
	second := now.Second()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	minute := now.Minute()
	pr("TOdays Date is ", month, day, year, hour, minute, second)
}