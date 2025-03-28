package main

import (
	"fmt"
	"os"
)

func main() {
	someMap := map[int]string{"305": "sue", "204": "Bob", "631" : "Jake", "073" : "tracy"}
	fmt.Println(someMap[os.Args[1]])
}