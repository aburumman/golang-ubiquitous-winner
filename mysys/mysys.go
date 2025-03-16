package mysys 

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
func GetHelloWorld()  string {
	return "Hello, World"
}
func GetUserID() int {
    return os.Getuid()
}