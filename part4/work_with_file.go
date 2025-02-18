package main

import (
	"fmt"
	//"os"
	"unicode/utf8"
)

func main() {

	fmt.Println(count_string([]string{"hello", "howare you", "?", "<<", ">>", "&"}))
}
func count_string(words []string) int {
	total := 0
	for _, word := range words {
		total += utf8.RuneCountInString(word)
	}
	total += len(words) -1
	
	return total
}

