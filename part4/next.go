package main

import (
	"fmt"
)

func somelist(mylist []string) []string {
	var nextlist []string
	nextlist = append(nextlist, "en_US", "fr_FR", "de_DE",)
	nextlist = append(nextlist, mylist...)

	return nextlist
}

func getPassedArg([]string) []string{
	var somelist []string

	myArgs := os.Args[]
}
func main() {}