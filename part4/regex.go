package main 

import (
	"fmt"
	"regexp"
	"time"
	"log"
)


func main() {
defer timeit("regexp")()
match, _ := regexp.MatchString("p([a-z]+)ch", "preach")

fmt.Println(match)

r, _ := regexp.Compile("p([a-z]+)ch")
fmt.Println(r.MatchString("pach"))
fmt.Println(r.FindString("peach"))
fmt.Println(r.FindStringIndex("preach"))
}

func timeit(name string) func(){
start := time.Now() 

return func () {
	taken := time.Since(start)
	log.Printf("Time taken: %s duration: %v", name, taken)
}
}
