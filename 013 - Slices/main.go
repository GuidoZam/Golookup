package main

import "fmt"

func main() {
	// You can skip the array size
	primes := []int{ 1, 2, 3, 5, 7, 11, 13 }
	
	// len gives the length of the slice
	arrayMax := len(primes)

	// These slices means the same, they take all the items in the array
	fmt.Println(primes[0:arrayMax])
	fmt.Println(primes[:7])
	fmt.Println(primes[0:])
	fmt.Println(primes[:])

	// Slice the slice to give it zero length
	s := primes[:0]
	printSlice(s)

	// Extend its length
	s = primes[:4]
	printSlice(s)

	// Drop its first two values
	s = primes[2:]
	printSlice(s)
}

func printSlice(s []int) {
	// len() gives the length of the slice
	// cap() gives the capacity of the slice and is the number of the elements in the underlying array
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}