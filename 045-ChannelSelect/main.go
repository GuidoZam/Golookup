package main

import (
	"fmt"
	"time"
)

func slowFunction(num int, c chan<- int) {
	result := num * 2
	time.Sleep(15 * time.Millisecond)
	c <- result
}

func fastFunction(num int, c chan<- int) {
	result := num * 2
	time.Sleep(5 * time.Millisecond)
	c <- result
}

func main() {
	slow := make(chan int)
	fast := make(chan int)

	go slowFunction(4, slow)
	go fastFunction(9, fast)

	select {
		case r := <- slow:
			fmt.Println("Slow finished first! Result: ", r)
		case r := <- fast:
			fmt.Println("Fast finished first! Result: ", r)
	}
}