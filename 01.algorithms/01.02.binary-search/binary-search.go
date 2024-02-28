package main

import (
	"fmt"
	"math"
)

// Data need to be sorted before using a binary search
// Worst case: O(log n)

func binarySearchInt(splice []int, target int) int {
	var midpoint float64
	var lowerBound int = 0
	var upperBound int = len(splice) - 1

	// Indefinite loop until the target is found
	// or the splice has been exhausted
	for {
		// Calculate the midpoint between the bounds,
		// rounding down where applicable
		midpoint = math.Floor((float64(lowerBound) + float64(upperBound)) / 2)

		// Check if the midpoint or bounds are the target
		if splice[int(midpoint)] == target {
			return int(midpoint)
		} else if splice[lowerBound] == target {
			return lowerBound
		} else if splice[upperBound] == target {
			return upperBound
		} else {
			// If the bounds are equal and the target still has not
			// been found, then the list has been exhausted
			if upperBound == lowerBound {
				return -1
			} else {
				// Adjust bounds accordingly before the next iteration
				if target < splice[int(midpoint)] {
					upperBound = int(midpoint) - 1
				} else {
					lowerBound = int(midpoint) + 1
				}
			}
		}
	}
}

func main() {
	// Search through a splice of integers
	intSet := []int{15, 29, 53, 99, 429, 500, 502, 505, 789, 911}
	intTarget := 29
	position := binarySearchInt(intSet, intTarget)
	if position != -1 {
		fmt.Println(intTarget, "is at position", position)
	} else {
		fmt.Println(intTarget, "was not found in the dataset")
	}
}
