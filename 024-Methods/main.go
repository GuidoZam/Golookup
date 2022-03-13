package main

import f "fmt"

type Person struct {
	FirstName, LastName string
}

// You can write a method in this way
func (p Person) GetFullName() string {
	return f.Sprintf("%v %v", p.FirstName, p.LastName)
}

// You can also write a method in this way
func GetInitials(p Person) string {
	return f.Sprintf("%v%v", p.FirstName[0:1], p.LastName[0:1])
}

func main() {
	v := Person{ "Jane", "Doe" }
	f.Println(v.GetFullName())
	f.Println(GetInitials(v))
}

