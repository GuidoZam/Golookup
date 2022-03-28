package main

import "fmt"

func classicForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func continuedForLoop() {
	seed := 1
	sum := 0
	for ; sum < 10;  {
		sum += seed
	}
	fmt.Println(sum)
}

func forLoopAsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteForLoop() {
	sum := 1
	for {
		sum += sum
		if (sum > 1000) {
			fmt.Println(sum)
			return
		}
	}
}

func main() {
	classicForLoop()
	continuedForLoop()
	forLoopAsWhile()
	infiniteForLoop()
}