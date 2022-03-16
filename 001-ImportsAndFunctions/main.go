package main

import (
	"fmt"
	m "math"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	// Write the sum of the two numbers
	fmt.Println(add(42, 13))

	// Write a constant from the imported package
	fmt.Println(m.Pi);
}