package main

import f "fmt"

type Person struct {
	FirstName, LastName string
}

func (p *Person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	p := Person{ "Jane", "Doe" }
	f.Println(p.GetFullName())
}