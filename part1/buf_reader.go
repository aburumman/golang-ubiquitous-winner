package main 

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Please tell me your name")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is %v", name)
}