package main

import "fmt"

func main() {
	// String concatenation
	fmt.Println("go" + "lang")

	// Integer operations
	// When using a comma, a space is added
	// + cannot be used in this instance, to avoid mismatched types
	fmt.Println("5 * 5 =", 5*5)
	fmt.Println("20 / 4 =", 20/4)

	// Float operations
	fmt.Println("22.50 / 4.00 =", 22.50/4.00)
	fmt.Println("3.33 * 3.33 =", 3.33*3.33)

	// Logical operations
	// True are false are spelt lowercase
	// || for OR, && for AND, ! to invert
	fmt.Println("true OR false =", true || false)
	fmt.Println("true AND false =", true && false)
	fmt.Println("true AND NOT false =", true && !false)
}
