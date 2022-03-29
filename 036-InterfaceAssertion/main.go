package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	var i interface{} = "this is a string"

	// cast i to string
	s := i.(string)
	fmt.Println(s)

	// test if casting i to string is possible
	s, ok := i.(string)
	fmt.Println(s, ok)

	// test if casting i to float64 is possible
	f, ok := i.(float64)
	fmt.Println(f, ok)

	var human interface{} = Person { Name: "Ezechiel" }
	p := human.(Person)
	fmt.Println(p)

	// this result in panic because i value is not of type Person
	// and this won't verify if casting is possibile
	// it would be better to use r, ok := human.(Person)
	r := i.(Person)
	fmt.Println(r)
}