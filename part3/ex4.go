package main 

import (
	"fmt"
	//"time"
	"os"
	"io"
	"net/http"
	//"ioutil"
)

func main() {
file, err := os.Create("./fromString.txt")
checkError(err)
content := "Jello world from go"
defer file.Close()
length, err := io.WriteString(file, content)
checkError(err)
fmt.Println("Wrote content: to file with lenght %v",  length)
readFile("./fromString.txt")
//myUrl:= "http://services.explorecalifornia.org/json/tours.php"
myUrl := "https://google.com" //"https://myf5.com"
fmt.Println(getUrl(myUrl))



}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(file string) {
	text, err := os.ReadFile(file)
	checkError(err)
	fmt.Println(string(text))
}

func getUrl(url string) string{
	response, err := http.Get(url)
	checkError(err)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	checkError(err)
	text := string(body)

	return text
}