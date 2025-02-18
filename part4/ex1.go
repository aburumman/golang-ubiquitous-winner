package main

import (
	"fmt"
	"net/http"
)

func main() {
server()


}

func server() {
	myServer := &http.Server{
		Addr: ":9000",
		Handler: http.HandlerFunc(basicHandler),
	}

	dist := myServer.ListenAndServe()
	if dist !=nil {
		fmt.Println(dist)
	}

}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Hello world")
	w.Write([]byte("Hello Wolrd"))
}