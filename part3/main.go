// package main  

// import (
// 	"fmt"
// 	"encoding/csv"
// 	"os"
// )

// func myfunc(file string, []item) error {

// 	new_file := os.Create(file)

// 	if err != nil {
// 		return err
// 		//os.Exit(1)
// 	}

// 	row := []string {"sku", "name"}

// 	writer := csv.NewWriter(new_file)
// 	defer write.flush()

// 	writer.Write(row)
// 	if err != nil {
// 		return err
// 	}

// }

// type item struct {
// 	sku string
// 	name string
// }

package main

import (
	"fmt"
	"encoding/csv"
	"os"
)

func main() {

	new_file, err := os.Create("file.csv")
	if err != nil {
		fmt.Println("%v", err)
	}

	writer := csv.NewWriter(new_file)
	defer writer.flush()
	row1 := []string{"sku", "name"}
	writer.Write(row1)
	if err != nil {
		fmt.Println("%v", err)
	}
}

type item struct {
	sku string
	name string
}