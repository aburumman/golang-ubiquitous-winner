package main 

import (
	"fmt"
)


const (
	successmsg = "success go"
	panicmsg = "panic go"
)

type Person struct {
	Name string 
	Age int
}

type Persons struct {
	m map[string]Person
}

func NewPersons() *Persons{
	return &Persons{
		m : make(map[string]Person),
	}
}

func (p *Persons) PanickingAdd(m Person) {
	if m.Age < 0 || m.Age > 100 {
		panic(fmt.Sprintf("Add: invalid age %d range for name: %s ", p.Age, p.Name))
	}
	p.m[m.Name] := m
	//p[m.Age] := m.Age
}

func PopulateData(data []Person) (result string) {
	result = successmsg
	defer func() {
		if r := recover(); r != nil {
			result = panicmsg
		}
	}()

	myp := NewPersons()
	for _, one := range data {
		myp.PanickingAdd(one)
	}

}