package main 

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	start_time := time.Now()
	go do_something()
	go doSomething()
	wg.Wait()
	fmt.Println(time.Since(start_time).Seconds())

}
func do_something() {
	time.Sleep(time.Second*2)
	fmt.Println("Completed something2")
	wg.Done()
}

func doSomething() {
	time.Sleep(time.Second*2)
	fmt.Println("completed somthing")
	wg.Done()
}