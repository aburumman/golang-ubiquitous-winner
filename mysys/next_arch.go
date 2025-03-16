package main 

import (
	"fmt"
	"os"
)

func main() {
	fileName := "locker.txt"
	content := "Tomorrow is gonna be a greate day"
	createFile(fileName)
	defer func () { err := closeFile(*fileName)
	if err != nil {
		fmt.Println(errr)
	}
} ()
	written := writeFile(*fileName, content)
	fmt.Println(written)

}

func createFile(fileName string) *os.File{
	newFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return newFile
}

func writeFile(file *os.File, content string)  string{

_, file_content := fmt.Println(content, file.Name())

return fmt.Sprintf("Content :", file_content, "written to %s", file)
}

func closeFile(file *os.File) error{
close_file := file.Close()

if close_File != nil {
return errors.New("Faild to close file", close_File)
}

return nil
}