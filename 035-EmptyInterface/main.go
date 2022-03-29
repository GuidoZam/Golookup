package main

import "fmt"

func main() {
	// empty interface
	var i interface{}
	describe(i)

	// every type is of type interface{}
	i = 42
	describe(i)

	// every type is of type interface{}
	i = "This is a string"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

