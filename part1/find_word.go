package main

import (
	"fmt"
	"strings"
)

func main() {
	words := `
	Needlses and pins
	Needles and pins
	sew me a sial to catch the wind
	`
	my_map := map[int]string{}
	word_list := strings.Fields(words)
	for index, word := range word_list {
		my_map[index] = word
	}
	fmt.Println(my_map)

}