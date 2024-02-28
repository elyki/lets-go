package main

import "fmt"

// For is Go's while

func main() {
	i := 0

	// Single condition loop
	for i <= 2 {
		fmt.Println("Iteration", i)
		// i = i + 1, or i += 1, or i++ all work
		i = i + 1
	}

	// Initial, condition, after loop
	for j := 0; j < 3; j++ {
		fmt.Println("Iteration", j)
	}

	// Range loop
	for i := range 3 {
		fmt.Println("Iteration", i)
	}

	// Indefinite loop
	for {
		fmt.Println("Iteration")
		break
	}

	// Iterations can be triggered with continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Iteration", n+1)
	}
}
