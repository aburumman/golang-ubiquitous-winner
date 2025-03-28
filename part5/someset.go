package main

import (
	"fmt"
	"os"
)

func getPassarg(minArgs int) []string {

	if len(os.Args) < minArgs {
		fmt.Printf("length of arguments must be greater than 1")
		os.Exit(1)
	}
	var args []string
	for i := 0; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args

}

func findLongeest(words []string) (string, error) {
	var longest string
	if len(words) <= 1 {
		return words[0], nil
	} else if len(words) <= 0 {
		return longest, fmt.Errorf("To can't supply an empty list")
	}
	for i, _ := range words {
		if len(words[i]) > len(longest) {
			longest = words[i] 
	}
}

return longest, nil
}

func main() {
	words := []string{"foo", "bar", "baz", "some gum", "lecture"}
	fmt.Println(findLongeest(words))
	fmt.Println(findLongeest(getPassarg(3)))
}