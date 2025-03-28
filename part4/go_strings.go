package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	firstName := "Mustapha"
	lastName := "Alaaynde"
	//rpeated_name := []string{"Mustapha", "Alayande", "Mustapha", "Alaayande"}
	rpeated_name := "Mustapha Alayande  Mustapha Alaaynde"
	compare := strings.Compare(firstName, lastName)
	fmt.Println(compare)
	lpare := strings.Compare(firstName, firstName)
	fmt.Println(lpare)
	//fmt.Println(strings.Contains(firstName, lastName))
	//fmt.Println(strings.Contains(rpeated_name, lastName))
	//fmt.Println(strings.Count(rpeated_name, lastName))
	//fmt.Println(strings.Split(rpeated_name, " "))
	//fmt.Println(strings.HasSuffix(rpeated_name, lastName))
	//fmt.Println(strings.HasPrefix(rpeated_name, firstName))
	//fmt.Println(strings.Index(rpeated_name, lastName))
	//fmt.Println(strings.IndexByte(rpeated_name))
	//fmt.Println(strings.Join([]string{firstName, lastName}, "-"))
	//fmt.Println(strings.Repeat( string(lastName +"-"), 5))
	//fmt.Println(strings.Replace(rpeated_name, "a", "x", -1))
	xs := strings.SplitAfter(rpeated_name, "h")
	forFunc(xs)
	fmt.Println(len(xs))
	fmt.Println(strings.ToLower(firstName))
	shi := "some PeoPLE"
	fmt.Println(strings.ToUpper(shi), strings.Title(shi))
	fmt.Println(strings.ToUpper(shi), strings.ToTitle(shi))
	fmt.Println(strings.ToLowerSpecial(unicode.AzeriCase, shi))
	fmt.Println(strings.TrimSuffix("Mustapha - koko", "koko"))
	fmt.Println(strings.TrimSpace("   This contains dome spcae   "))

	//fmt.Println(strings.SplitAfter(rpeated_name, "h"))
	//fmt.Println(strings.SplitN(rpeated_name, "h", 5))

}

func forFunc(name []string)  {
	for _, x := range name {
		fmt.Println(x)
	}
}