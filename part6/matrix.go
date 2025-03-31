package main 

import (
	"fmt"
)

func main() {
	var text [3][3]int = [3][3]int {[3]int{1, 2, 3}, [3]int{4,5,6}, [3]int{7,8,9}}
	text[0] = [3]int{1, 2, 3}
	text[1] = [3]int{34, 56,78}
	b := [3]int{}
	//x := copy(b, text[2])
	fmt.Printf( "%v,",  text)
	//fmt.Println(b, x)
	xy := []int{}
	c := []int{1, 2, 3}
	//xy = append(xy, c[:])
	fmt.Println(xy)
}