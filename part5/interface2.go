package main 

import (
	"fmt"
)

type Employee struct {
	BasePerson 
	Salary int 
	LineManager *Manager
}
type Manager struct {
	Employee
}
type BasePerson struct {
	FirstName string
	LastName string
}
type Person interface {
	GetName() string
}

func (p *Employee) GetName() string {
	return p.FirstName
}

func (p *Manager) GetName() string {
	return p.FirstName
}

func SayHello(p Person) {
	fmt.Println(p.GetName())
}

func main() {
	tom := &Manager{
		Employee: Employee{
			Salary:  6000,
			LineManager: nil,
		BasePerson: BasePerson{
			FirstName: "Tomiwa",
			LastName: "Shikeni",
		},
	}}

	ken := &Employee {
		BasePerson: BasePerson{
			FirstName: "Kenneth",
			LastName: "Werty",
		},
		Salary: 9_000,
		LineManager: tom,
	}
	SayHello(ken)
	SayHello(tom)
}

// 	myp := BasePerson {
// 	FirstName: "Tomiwa",
// 	LastName: "Shikeni",
// }