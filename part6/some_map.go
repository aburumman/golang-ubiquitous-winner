package main 

import (
	"fmt"
)

func funcy() {
	var i interface{} = nil

	switch i.(type) {
	case bool:
		fmt.Println("boolean")
	case int:
		fmt.Println("int")
	case int64:
		fmt.Println("int64")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Printf("\n %T \n", i)
	}
}
func main() {
	funcy()
	myMap := map[string]string{"Lagos": "Ikeja", "Niger": "Minna"}
	fmt.Println(myMap["Lagos"])
	delete(myMap, "Niger")
	some_state, ok := myMap["Oguyn"]
	fmt.Println(some_state, ok)

	type Doctor struct {
		name string
		number int
		companion []string
	}
	doctor1 := Doctor{
		name: "Shackles",
		number: 70,
		companion : []string{"Exploedi", "Ickle", "Nextremely"},
	}
	fmt.Println(doctor1.companion[0])
	doctor2 := doctor1
	doctor2.number = 890
	fmt.Println(doctor2)

	type Animal struct {
		Origin string `default: "Earth"`
		Name string
		Legs int
	}
	type Bird struct {
		Animal
		CanFly bool
		Colors []string
	}
	ost := Bird{CanFly: false, Colors: []string{""}, Animal: Animal{Origin: "Kenya", Name: "Ostrich", Legs: 2},}
	fmt.Printf("%t, %T", ost, []byte("Man"))

	if exist, ok := myMap["Lagos"]; ok {
		fmt.Println("\n", exist)
	}
}