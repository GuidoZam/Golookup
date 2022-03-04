package main

import "fmt"

func main() {
	var stringArray [2]string
	stringArray[0] = "Hello"
	stringArray[1] = "World"
	fmt.Println(stringArray[0], stringArray[1])
	fmt.Println(stringArray)

	fibonacciArray := [9]int{ 1, 1, 2, 3, 5, 8, 13, 21, 34 }
	fmt.Println(fibonacciArray)

	// Slice an array
	var partial []int = fibonacciArray[5:8]
	fmt.Println(partial)

	// Slice array and update value
	colors := [5]string{
		"Red",
		"Blue",
		"Green",
		"Yellow",
		"Purple"}
	fmt.Println(colors)

	// Get the first and second colors
	a := colors[0:2]
	// Get the second and third colors
	b := colors[1:3]
	// Get the fourth and fifth colors
	c := colors[3:5]

	fmt.Println(a, b, c)

	// Update the value of the sliced array updates also the value of the parent array 
	b[0] = "XXXXX"
	fmt.Println(a, b, c)
	fmt.Println(colors)
}