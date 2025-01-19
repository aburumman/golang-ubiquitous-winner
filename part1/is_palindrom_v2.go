package main

import (
	"fmt"
	"strings"
)


func is_palindrom(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	start, end := 0, len(s)-1
	
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func main() {
	words := []string{"civic", "madam", "turnup","radar", "kayak", "redivider", "razor","racecar", "reviver", "printer","deified", "rotator", "repaper","excellent"}
	for index, word := range words {
		if (is_palindrom(word)) {
			fmt.Println("%s %s is a palindrom", index, word)
		} else {
			fmt.Println("%s %s is not a palindrom", index, word)
		}
		//fmt.Println(is_palindrom(word))
	}
	//fmt.Println(is_palindrom(""))
}