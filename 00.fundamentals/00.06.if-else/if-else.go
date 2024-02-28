package main

import "fmt"

const num int = 21

func main() {
	// Standard if/else
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// If/else if/else
	if num%6 == 0 {
		fmt.Println(num, "is a multiple of 6")
	} else if num%7 == 0 {
		fmt.Println(num, "is a multiple of 7")
	} else {
		fmt.Println(num, "is neither a multiple of 6 nor 7")
	}

	// Logical operations
	if num%3 == 0 && num%7 == 0 {
		fmt.Println(num, "is a multiple of 3 and 7")
	}

	// Variables can be declared prior to a conditional
	if n := 0; n > 0 {
		fmt.Println(n, "is greater than 0")
	} else {
		fmt.Println(n, "is not greater than 0")
	}
}
