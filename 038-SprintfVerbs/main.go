package main

import "fmt"

func main() {
	numericValue := 42
	stringValue := "Hello world!"
	booleanValue := true
	helloString := "hello"
	worldString := "world"

	// %v -> generic verb
	fmt.Printf("%v %v %v\n", numericValue, stringValue, booleanValue)

	// %s -> string verb
	fmt.Printf("%s %s %s\n", numericValue, stringValue, booleanValue)

	// %q -> quote verb
	fmt.Printf("%q %q %q\n", numericValue, stringValue, booleanValue)

	// %d -> base 10 decimal verb
	fmt.Printf("%d %d %d\n", numericValue, stringValue, booleanValue)

	// %t -> boolean verb
	fmt.Printf("%t %t %t\n", numericValue, stringValue, booleanValue)

	// Argument indexing 1-based
	fmt.Printf("%[1]s %[2]s\n", helloString, worldString)
	fmt.Printf("%[1]s %[1]s\n", helloString)
	fmt.Printf("%[1]s %[1]q\n", worldString)
	fmt.Printf("%[1]s %[2]q\n", helloString, worldString)
}