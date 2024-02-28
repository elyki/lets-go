package main

import "fmt"

// Parameters and the output must have types
func add(num1 int, num2 int) int {
	return num1 + num2
}

// Recursion can be used, too
func factorial(num int) int {
	if num > 0 {
		return num * factorial(num-1)
	} else {
		return 1
	}
}

func main() {
	fmt.Println(add(1, 4))
	fmt.Println(add(add(1, 4), 2))

	fmt.Println(factorial(5))
}
