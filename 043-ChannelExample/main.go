package main

import (
	"fmt"
	"time"
)

func multiply(a int, b int, out chan<- int) {
	fmt.Print("Calculating...")
	result := a * b

	time.Sleep(2000 * time.Millisecond)

	out <- result
	fmt.Println("done!")
}

func main() {
	a := 3
	b := 6

	result := make(chan int)

	go multiply(a, b, result)

	fmt.Println(<- result)
}