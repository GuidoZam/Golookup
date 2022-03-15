package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   		bool       	= false
	sampleText	string	   	= "This is just an example"
	MaxInt 		uint64     	= 1<<64 - 1
	sampleInt 	int     	= 42
	z      		complex128 	= cmplx.Sqrt(-5 + 12i)
	a			rune		= 12312
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", sampleText, sampleText)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", sampleInt, sampleInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", a, a)
}