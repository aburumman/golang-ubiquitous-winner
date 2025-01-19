package main 

import (
	"fmt"
	//"sync"
)

func main() {
	data_chan := make(chan string)

	my_name := "Mustapha"
	data_chan <- my_name
	next_name := ""
	next_name = <- data_chan
	//go data_chan <- my_name
	go func () { 
		fmt.Printf(next_name) 
	} ()
	go fmt.Println(my_name, data_chan, )
}