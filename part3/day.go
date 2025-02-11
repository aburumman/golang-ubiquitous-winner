package main   

import (
	"fmt"
)

func main() {
	var enter_day string
	day_prompt, _ := fmt.Println("Enter between 1 - 7 to get day of the week")
	day, err := fmt.Scanln(&enter_day)
	if err != nil {
		return "Please choose a number between 1 - 7"
}
}

switch day = enter_day {
case 1:
	return fmt.Println("Sunday")
case 2:
	return fmt.Println("Monday")
case 3:
	return fmt.Println("Tuesday")
case 4:
	return fmt.Println("Wednesay")

}