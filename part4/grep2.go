package main 


import (
	"fmt"
	"strings"
	"bufio"
	"io"
)

func main() {
	word_stream := strings.NewReader("THis the word\n, I want to say the word\n, but can say nothing\n, so i need to try the word")
	fmt.Println(grep(word_stream, "say"))
}

func grep(text io.Reader, word string) ([]string, error){
   matches := []string{}
   texts:= bufio.NewScanner(text)
   //errCheck()
   for texts.Scan() {
	if strings.Contains(texts.Text(), word) {
		matches = append(matches, texts.Text())
	}
	
   }
   err := texts.Err()
   if err !=nil {
	return nil, err
   }
   return matches, nil

}

// func errCheck() {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }