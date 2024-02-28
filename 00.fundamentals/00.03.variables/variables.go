package main

import "fmt"

func main() {
	// Declaring one variable
	var name string = "seb"
	fmt.Println("Name:", name)

	// Declaring multiple variables
	var num1, num2 int = 5, 7
	fmt.Println(num1, "+", num2, "=", num1+num2)

	// Declaring variables of different types
	// := can be used to achieve this
	// Go can infer the type of variables
	num3, label := 9, "Number 3"
	fmt.Println(label, ":", num3)
}
