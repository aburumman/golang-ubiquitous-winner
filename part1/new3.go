package main 

import (
	"fmt"
)


func main() {
	fmt.Println(isOdd(25))
	values := []int{1,2,3,4,5,6,7,8}
	//make_filter := filter(isOdd, values )
	filter(isOdd, values )
	//fmt.Println(make_filter)
}

func filter(pred func(int) bool, myslice []int) []int {
	//return_list := make([]int, 0, len(myslice)) // out := []int{}
	var return_list []int
	for _, v := range myslice {
		if pred(v){
			return_list = append(return_list, v)
			//return_list[x] = v
		}
	}
 return return_list
}

func isOdd(x int) bool{
	if x % 2 == 1 {
		return true
	}
	return false
}

