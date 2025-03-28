package main  

import (
	"fmt"
	"errors"
)

func main() {

	fmt.Println(calcMean([]int{ 56, 23, 12, 80, 9,}))

}

func calcMean(number []int) (*float64, error) {
	if len(number) <= 1 {
		//return nil, fmt.Errorf("Only one values proived'")
		return nil, errors.New("Onyl one value provided")
	}
	var total int
		for _, num := range number{
			total += num
		}
	mean := float64(total) / float64(len(number))
	
	return &mean, nil
}