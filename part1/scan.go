package main

import "fmt"

func scanwithscan() {
	var d1, d2 int 
	var s1 string 

	fmt.Println("Scanning input")
	if _, err := fmt.Sscan("5 7 \n9", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return 
}
fmt.Println(d1, d2, s1)
}

func main() {
	scanwithscan()
}