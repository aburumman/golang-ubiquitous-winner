package main

import (
	"fmt"
	"strings"
)

func main() {
	
	fmt.Println(fold("A"))
}

type letter struct{
	english string 
	arabic string
}

//var alif = []letter { {"A", "l"},}
var alif = letter{"A", "l"}



func fold(lang string) (string, error) {
	for _, x := range alif {
		if (strings.EqualFold(lang, x.english)) {
			return x.arabic, nil
		}
	}
	return "", fmt.Errorf("Unknow arabic letter", lang) 
}