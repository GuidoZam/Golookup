package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	BirthCity string
}

func main() {
	p1 := Person{ "John", "Doe", "" }
	fmt.Println(p1)
	p1.BirthCity = "New Orleans"
	fmt.Println(p1)

	// Pointer with struct
	p2 := &p1
	p2.FirstName = "Jane"
	fmt.Println(*p2)
}