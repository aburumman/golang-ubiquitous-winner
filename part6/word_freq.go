package main

import (
	"fmt"
	"math/rand"
)

func word_freq(words []string)  map[string]int {
	word_cloud := make(map[string]int)
	for _, word := range words {
		word_cloud[word]++
	}

	return word_cloud
}


func rol_dice() int {
return rand.Intn(7)
}