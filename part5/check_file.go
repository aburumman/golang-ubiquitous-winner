package main 

import (
	"fmt"
	"os"
)

func checkFile(file string) (bool, error) {
	// Check if a file exists
	fileName := file 

	_, err := os.Stat(fileName)
	if err != nil {
		return false, err
	}

	return true, nil
}


func writeFile(filename, content string) error {
	written_file, error := os.WriteString(fileName, content)
}

// func main() {
// 	fmt.Println("Please enter the full path and name of file")
// 	var theFile string 
// 	fmt.Scanln(&theFile)
// 	fmt.Println("Your file name is:", theFile )
// 	fmt.Println(checkFile(theFile))
// }