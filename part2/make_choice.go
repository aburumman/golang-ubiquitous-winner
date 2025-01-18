package main

import ("fmt")

func main() {

	var num int 

	fmt.Println("Which chapter would you like to move to")
	fmt.Scanf("%d", &num)
	//num_choice, err := fmt.Scanf("%d", &num)
	// if err != nil {
	// //fmt.Println(err)
	// //num_choice = 10
	// num = 10
	// }

	switch num {
	case 1: gotoChapter(1)
	case 2: gotoChapter(2)
	case 3: gotoChapter(3)
	case 4: gotoChapter(4)
	default: gotoChapter(num)
	}
}

func gotoChapter(chapter int)  {
		fmt.Println("This Chapter %v", chapter)
}