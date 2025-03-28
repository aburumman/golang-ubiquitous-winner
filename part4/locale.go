package main  


import (
	"fmt"
	"os"
)

func getLocal() ([]string){
	local := []string{"en_US", "en_GB", "fr_FR", "en_CA"}

	add_local := getArgs()
	for i:= 0; i < len(add_local); i++ {
		local = append(local, add_local[i])
	}

	return local
}

func getArgs() []string{
	var lang []string
	if len(os.Args) < 1 {
		fmt.Println("You most provide at least one language")
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		lang = append(lang, os.Args[i])
	}
	return lang
}


func findLongest(words []string) string {
	var longest string
	for i, word :=range words {
		if len(word) > len(longest) {
			longest = words[i]
	}
		} 
		return longest
}

func main() {
	fmt.Println(getLocal())
	mylist := []string{"lets","somewords", "go", "home", "bcause", "tomorrow", "still the same"}
	fmt.Println(findLongest(mylist))
}