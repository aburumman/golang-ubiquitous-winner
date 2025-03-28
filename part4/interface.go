package main  

import (
	"fmt"
)

type Person interface{
	GetName() string
}
type BasePerson struct {
	First string 
	Last string 
}

type Employee struct {
	BasePerson
	Salary int 
	LineManager  *Manager
}

type Manager struct { 
	Employee
}

func (m *Manager) GetName() string{
	return m.First
}

func (m *Employee) GetName() string{
	return m.First
}

func sayHello(p Person) {
	fmt.Println(p.GetName())
}


func main() {

	fManager := &Manager{
		Employee: Employee{
			LineManager: nil,
			Salary: 60_000,
			BasePerson: BasePerson{
				First: "Jide",
				Last: "Onuora",
			},
		},
	}
	fEmployee:= &Employee{
		BasePerson: BasePerson{
		First: "Segundo",
		Last: "Onuora",
		},
		Salary: 30_000,
		LineManager: fManager,
	}

	sayHello(fEmployee)
	sayHello(fManager)
}
