package main

import "fmt"

func say(c chan string) {
	text := <- c
	fmt.Println(text)
}

func ping(p chan string) {
	p <- "ping!"
}

func main() {
	c := make(chan string)

	go say(c)
	c <- "Hi!"

	go say(c)
	c <- "Hej!"
	
	pingChannel := make(chan string)

	go ping(pingChannel)

	pingResult := <- pingChannel
	fmt.Println(pingResult)
}