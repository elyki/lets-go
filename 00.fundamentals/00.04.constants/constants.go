package main

import "fmt"

// const is used to declare a constant
const pi float32 = 3.14

func main() {
	var r int = 5.00

	// Types need to be converted to avoid type mismatch
	fmt.Println("Area:", pi*(float32(r)*float32(r)))
}
