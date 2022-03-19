package main

import (
	"fmt"
)

type Person struct {
	FirstName, LastName string
}

func (p *Person) GetFullName() {
	fmt.Println(p.FirstName, p.LastName)
}

func GetFullNameFunc(p *Person) {
	fmt.Println(p.FirstName, p.LastName)
}

func main() {
	p := Person { "Jane", "Doe" }
	p.GetFullName()
	GetFullNameFunc(&p)

	p2 := &Person{ "John", "Doe" }
	p2.GetFullName()
	GetFullNameFunc(p2)
}