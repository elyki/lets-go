package main

import "fmt"

// Data do not need to be sorted before using a linear search
// Worst case: O(n), suited to smaller datasets

func linearSearchInt(splice []int, target int) int {
	// Iterate through every item in the splice
	for index, value := range splice {
		// Return the index if found
		if value == target {
			return index
		}
	}
	// After splice has been exhausted, return -1
	return -1
}

func linearSearchString(splice []string, target string) int {
	// Iterate through every item in the splice
	for index, value := range splice {
		// Return the index if found
		if value == target {
			return index
		}
	}
	// After splice has been exhausted, return -1
	return -1
}

func main() {
	// Search through a splice of integers
	intSet := []int{9, 2, 4, 1, 3, 8, 0, 7, 6, 5}
	intTarget := 5
	position := linearSearchInt(intSet, intTarget)
	if position != -1 {
		fmt.Println(intTarget, "is at position", position)
	} else {
		fmt.Println(intTarget, "was not found in the dataset")
	}

	// Search through a splice of strings
	stringSet := []string{"cow", "dog", "crocodile", "elephant", "hippo", "quokka"}
	stringTarget := "elephant"
	position = linearSearchString(stringSet, stringTarget)
	if position != -1 {
		fmt.Println(stringTarget, "is at position", position)
	} else {
		fmt.Println(stringTarget, "was not found in the dataset")
	}
}
