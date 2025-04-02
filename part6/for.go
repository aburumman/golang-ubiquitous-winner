package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for ; i < 20; i++ {
		fmt.Println(i)
		i += 2
	}

	for i < 30 {
		fmt.Println("i is",i)
		fmt.Println(os.Getenv("HOME"))
		 i *=2
		 fmt.Println(os.Getenv("USER"))
		 fmt.Println("then ",i)

	}
}
