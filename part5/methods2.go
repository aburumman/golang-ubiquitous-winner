package main 

import (
	"time"
	"fmt"
)


func main() {
	tom := NewPerson("Tomiwa", "Mustapha", 1994, 4, 26)
	fmt.Println(tom.sayHello())
	fmt.Println(tom.GetAge())
}

type Person struct {
	First, Last string
	//Age  int 
	DOB  time.Time
}

func NewPerson (first, last string, year, month, day int)  (person *Person) {

	return &Person {
		First: first,
		Last: last,
		DOB: time.Date(year, time.Month(month), day, 0,0,0,0, time.Local),
	}

}

func (person Person) sayHello() string {
	return fmt.Sprintf("Hello %s", person.First)
}

func (person Person) GetAge() int {
	Age := time.Now().Year() - person.DOB.Year()
	return Age
}