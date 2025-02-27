package main 

import (
	"io"
	"strings"
	"fmt"
	"bufio"
	"os"
)

func main() {
	//t := os.Stdin
	fmt.Println("What's your name?")
	t, err := readsome(os.Stdin)
	fmt.Println(t, err)
	list_sentences := []string{"This isa a place for you And learn this also", "if you want to know more Let's go then for you to see the", "THere's more we can do here" }
	//text2 := strings.NewReader("This isa a place for you And learn this also, if you want to know more Let's go then for you to see the")
	var sentences []string
	for _, sent := range list_sentences {
		out, err := grep(strings.NewReader(sent), "you")
		if err == nil {
			sentences = append(sentences, out[0])
		}
	}
	fmt.Println(sentences)
	//out, err := grep(text2, "you")
	//fmt.Println(out, err)

}


func readsome (text *os.File) (string, error) {
  mytext := bufio.NewScanner(os.Stdin)
  if mytext.Scan() {
	return mytext.Text(), nil
  }
  err := mytext.Err()
  return mytext.Text(), err
}

func grep (in io.Reader, str string) ([]string, error) {
	var matches []string
	s := bufio.NewScanner(in)
	for s.Scan() { 
			if strings.Contains(s.Text(), str) {
				matches = append(matches, s.Text())
			}
	}
	
	return matches, nil
}