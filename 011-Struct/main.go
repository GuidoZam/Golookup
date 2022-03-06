package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	BirthCity string
}

type Triangle struct {
	a int
	b int
	c int
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

	// Constructor without values
	t := Triangle {}
	fmt.Println(t)

	// Construct with partial variables
	p3 := Person { FirstName: "Luigi" }
	fmt.Println(p3)

	p4 := Person { FirstName: "Mario", LastName: "Super" }
	fmt.Println(p4)

	p5 := Person { BirthCity: "Berlin" }
	fmt.Println(p5)

}