package main

import "fmt"

func main() {
	f := 42.9
	fmt.Printf("f with value %v is of type %T\n", f, f)

	i := 7
	fmt.Printf("i with value %v is of type %T\n", i, i)
	
	t := "test"
	fmt.Printf("t with value %v is of type %T\n", t, t)
}