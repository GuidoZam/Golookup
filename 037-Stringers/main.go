package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Hi! I'm %v and I am %v years old", p.Name, p.Age)
}

type Animal struct {
	Type string
	Color string
}

func (a Animal) String() string {
	return fmt.Sprintf("I'm a %v %v", a.Color, a.Type)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a)
	fmt.Println(z)

	blackCat := Animal{"cat", "black"}
	whiteDog := Animal{"dog", "white"}
	fmt.Println(blackCat)
	fmt.Println(whiteDog)
}