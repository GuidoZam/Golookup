package main

import (
	"fmt"
)

type Animal interface {
    Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
    return "Meow!"
}

type Duck struct {
}

func (d Duck) Speak() string {
    return "Quack!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Duck{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}

