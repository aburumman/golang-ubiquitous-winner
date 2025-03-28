package main

import (
	"fmt"
	"runtime"
	"log"
)

func main() {
	fmt.Println(runtime.Version())
	log.Printf(runtime.Version())
}