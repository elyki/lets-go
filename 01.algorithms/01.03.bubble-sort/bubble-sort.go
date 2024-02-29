package main

import "fmt"

// Swaps unordered, adjacent values
// Worst case: O(n^2)

func bubbleSortInt(splice []int) []int {
	var swapMade bool = true
	var buffer int

	// Repeats until a swap has not been made
	// i.e., the splice has been sorted
	for swapMade {
		swapMade = false
		// If adjacent values are unordered, swap them and reset
		// swapMade to true so another iteration occurs
		for i := 0; i < len(splice)-1; i++ {
			if splice[i] > splice[i+1] {
				buffer = splice[i]
				splice[i] = splice[i+1]
				splice[i+1] = buffer
				swapMade = true
			}
		}
	}

	// Return the sorted splice of integers
	return splice
}

// Adapted for bubble sorting a splice of strings
func bubbleSortString(splice []string) []string {
	var swapMade bool = true
	var buffer string

	for swapMade {
		swapMade = false
		for i := 0; i < len(splice)-1; i++ {
			if splice[i] > splice[i+1] {
				buffer = splice[i]
				splice[i] = splice[i+1]
				splice[i+1] = buffer
				swapMade = true
			}
		}
	}

	return splice
}

func main() {
	intSet := []int{1, 3, 2, 99, 21, 4}
	fmt.Println(bubbleSortInt(intSet))

	stringSet := []string{"cow", "armadillo", "hippo", "crocodile", "quokka", "elephant"}
	fmt.Println(bubbleSortString(stringSet))
}
