package main 

import (
	"fmt"
	"sort"
)


func main() {
	// sum mean median
	myList := []float64{10, 1, 2, 3}
 fmt.Println(sum(myList))
 fmt.Println(mean(myList))

}

func sum(numbers []float64) float64 {
	var start float64 = 0
	for _, num := range numbers {
		start += num
	}
	return float64(start)
}

func mean(numbers []float64) float64 {
lent := len(numbers)
//make_list := []float64
//list_sum := sum(make_list)
// if copy(make_list, numbers) == err {
// 	fmt.Println(err)
// }
myMean := sum(numbers) / float64(lent)
//median_2 := sum(make_list)

return myMean
}

func median(numbers []float64) float64 {
	var middle int
	myNumbers := make([]float64, len(numbers))
	//myNumbers = append(myNumbers, numbers[:len(numbers)-1])
	myNumbers =numbers[:len(numbers)-1]
	middle = len(myNumbers) / 2
	if len(myNumbers) % 2 != 0 {
return numbers[middle+1]
	} else {
the_median := (numbers[middle+1] + numbers[middle]) / 2
return the_median
	}
}