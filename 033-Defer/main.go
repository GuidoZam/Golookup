package main

import "fmt"

func deferredHelloWorld() {
	defer fmt.Println(" world!")

	fmt.Print("hello")
}

func deferredCounting() {
	fmt.Print("counting...")

	defer fmt.Println("LAUNCH!")

	for i := 1; i < 6; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done!")
}

func main() {
	deferredHelloWorld()
	deferredCounting()
}