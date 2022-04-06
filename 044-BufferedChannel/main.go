package main

import "fmt"

func main() {
	channel := make(chan int, 3)

	channel <- 4
	channel <- 9
	channel <- 13

	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println(<- channel)
}