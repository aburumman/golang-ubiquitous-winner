package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
myfile, err  := os.Open("learn.txt")
if err != nil {
	panic(err)
}
defer myfile.Close()
file_byte, err := readsome(myfile)
if (err != nil) {
	panic(err)
}
fmt.Printf("The number of bytes are %d", file_byte)
myText := strings.NewReader("Hello Dudumuffin")

text_out, _ := readsome(myText)

fmt.Printf("The number of bytes are %d", text_out)



}

func readsome(r io.Reader) (int, error) {
	count := 0 
	mybytes := make([]byte, 1024)


	//for {
	in, err := r.Read(mybytes)
		for _, x := range mybytes[:in] {
			if (x >= '0' && x <= '9') || (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z'){
			count++
		}
	}
	// if err == io.EOF {
	// 	return count, nil
	// }
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return 0, err
	}
//} 
return count, nil
}