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
	GetName()
}

func (p *Employee) GetName() string {
	return p.FirstName
}

func (p *Manager) GetName() string {
	return p.FirstName
}

func SayHello(p Person) {
	fmt.Println("Hello to you %s", p.GetName())
}

func main() {
	tom := &Manager{
		Employee{
			Salary:  6000,
			LineManager: nil,
		BasePerson: BasePerson{
			FirstName: "Tomiwa",
			LastName: "Shikeni",
		},
	}}

myp := BasePerson {
	FirstName: "Tomiwa",
	LastName: "Shikeni",
}