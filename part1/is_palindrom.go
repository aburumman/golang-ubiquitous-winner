package main

import (
	"fmt"
)

var word string 

func main () {
fmt.Printf("%s", checkword("aAddaa"))

}

func checkword(word string) []string {
	word_len := len(word)

	//if word_len % 2 ==  0 {
		word_part := []string{word[:word_len/2]}
		word_part2 := []string{word[word_len/2:len(word)]}
		bck := len(word_part)
		for _, j := range word_part {
			bck = bck -1
			if j == word_part2[bck] {
				continue
		} else {
			fmt.Println("%v is not a palindrome")
		}
		//return word_part, word_part2
		fmt.Println("%+v, %+v", word_part, word_part2)
	}
//}
return []string{word}
}