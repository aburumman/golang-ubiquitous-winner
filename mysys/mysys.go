package main  

import (

	//"golang.org/x/sys/unix"
	"fmt"
	"os"
	//"unix"
)

func main() {
	fmt.Println(GetUserID())
}

// func GetUserID() int {
// 	return unix.Getuid()

// }

func GetUserID() int {
    return os.Getuid()
}