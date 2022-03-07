package main

import f "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	f.Println("The value:", m["Answer"])

	// Getting the value also returns a variable that specify the existence of the element
	v, present := m["Answer"]
	f.Println("The value:", v, "Present?", present)

	m["Answer"] = 48
	f.Println("The value:", m["Answer"])

	delete(m, "Answer")
	f.Println("The value:", m["Answer"])

	v, present = m["Answer"]
	f.Println("The value:", v, "Present?", present)
}