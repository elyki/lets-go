package main

import "fmt"

func main() {
	// Create a static array of integers with a capacity of 5
	var aInt [5]int
	fmt.Println(aInt)

	// Update elements of an array
	aInt[4] = 100
	fmt.Println(aInt)

	// Fetch an element of an array
	fmt.Println(aInt[4])

	// Get the length of an array
	fmt.Println("Length of array:", len(aInt))

	// Arrays can be 2D
	var a2d [2][2]int
	fmt.Println(len(a2d))
	for i := range 2 {
		for j := range 2 {
			a2d[i][j] = i + i + j
		}
	}
	fmt.Println(a2d)
}
