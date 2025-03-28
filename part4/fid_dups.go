package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	myFile, err := parseFile("testfile.txt")
	if err !=  nil {
		panic(err)
	}
	out_file, err := get_files(myFile)
	fmt.Println(out_file)


}

func parseFile(filename string) ([]string, error) {
	newList := []string{}
	if _,fileX := os.Stat(filename); fileX != nil {
		return newList, fileX
	}
	theFile, _ := os.Open(filename)
	defer theFile.Close()
	scanner := bufio.NewScanner(theFile)
	for scanner.Scan() {
		line := scanner.Text()
		newList = append(newList, line)
	}

	return newList, nil


}
func get_files(someList []string) ([]string, error) {

	if len(someList) <= 1 {
		return someList, fmt.Errorf("file contains only one line")
	}
	var newList []string
	var seen = make(map[string]bool)
	for _, line := range someList {
		if seen[line] == false {
			newList = append(newList, line)
			seen[line] = true			
	}
}
return newList, nil
}