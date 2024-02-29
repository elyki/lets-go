package main

import "fmt"

// Moves unsorted values into the sorted part of a splice
// Worst case: O(n^2)

func insertionSortInt(splice []int) []int {
	var toInsert int
	var c int

	// Iterate through the inputted splice
	for i := range splice {
		// Mark the next element as for insertion
		toInsert = splice[i]
		c = i

		// Move leftwards along the sorted part of the splice
		// until the correct position is located
		for c > 0 && splice[c-1] > toInsert {
			splice[c] = splice[c-1]
			c--
		}

		// Re-insert the element in the correct position
		splice[c] = toInsert
	}

	// Return the sorted splice of integers
	return splice
}

// Adapted for insertion sorting a splice of strings
func insertionSortString(splice []string) []string {
	var toInsert string
	var c int

	for i := range splice {
		toInsert = splice[i]
		c = i

		for c > 0 && splice[c-1] > toInsert {
			splice[c] = splice[c-1]
			c--
		}

		splice[c] = toInsert
	}

	return splice
}

func main() {
	intSet := []int{1, 3, 2, 99, 21, 4}
	fmt.Println(insertionSortInt(intSet))

	stringSet := []string{"cow", "armadillo", "hippo", "crocodile", "quokka", "elephant"}
	fmt.Println(insertionSortString(stringSet))
}
