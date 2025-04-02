package main

import (
	"fmt"
)

func main() {
var someWork Writer = &ConsoleWriter{}
fmt.Println(someWork.Write([]byte("What is the real thing")))

}

type Writer interface {
	Write([]byte) (string, error)
}

type ConsoleWriter struct {}
func (cw ConsoleWriter) Write( data[]byte) (string, error) {
	return fmt.Sprintf(string(data)), nil
}
