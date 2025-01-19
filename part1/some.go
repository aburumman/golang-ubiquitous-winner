package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {
  fmt.Println("What is your name please: ")
  input := bufio.NewReader(os.Stdin)
  user_name, err  := input.ReadString('\n')
  if err != nil {
	log.Fatal(err)
	fmt.Println(err)
  } else {
	//fmt.Sprintf("Your name is %s", user_name)
	fmt.Println("YOur name is  :", user_name)
  }
}