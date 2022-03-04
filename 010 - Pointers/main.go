package main

import "fmt"

func main() {
	n := 38

	// set the pointer
	p := &n
	// read the pointer value
	fmt.Println(*p)
	// set the value through the pointer
	*p = 5
	fmt.Println(n)

	j := 3600
	p = &j
	*p = *p / 60
	fmt.Println(j)

	a := [3]string{ "a", "b", "c" }
	b := &a;
	fmt.Println(*b)

}