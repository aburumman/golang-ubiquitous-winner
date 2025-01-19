package main

import (
	"fmt"
	"net/http"
)

func get_cont(url string) (string, error){
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	content_type := response.Header.Get("Content-Type")

	if content_type == "" {
		return "", fmt.Errorf("no content header")
	}

	return content_type, nil
}

func main() {
	url := "https://youtube.com"
	my_header, err := get_cont(url)

	if err != nil {
		fmt.Println(err)
	} else {

	fmt.Println(my_header)
	}
}