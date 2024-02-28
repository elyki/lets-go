package main

import "fmt"

func main() {
	var num int = 1

	// default is the equivalent of else
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("I don't know how to spell that number")
	}

	// Multiple cases can be checked
	switch num {
	case 1, 2:
		fmt.Println("The number is either one or two")
	case 3, 4:
		fmt.Println("The number is either three or four")
	}
}
