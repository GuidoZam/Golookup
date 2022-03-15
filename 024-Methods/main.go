package main

import (
	f "fmt"
	"math"
)


type Person struct {
	FirstName, LastName string
}

// You can write a method in this way
func (p Person) GetFullName() string {
	return f.Sprintf("%v %v", p.FirstName, p.LastName)
}

// You can also write a method in this way
func GetInitials(p Person) string {
	return f.Sprintf("%v%v", p.FirstName[0:1], p.LastName[0:1])
}

// You can also define method on a non structured type
type MyFloat float64

func (floatVar MyFloat) Abs() float64 {
	if floatVar < 0 {
		return float64(-floatVar)
	}
	return float64(floatVar)
}

type MyInt int

func (intVar MyInt) Abs() int {
	if intVar < 0 {
		return int(-intVar)
	}
	return int(intVar)
}

func main() {
	v := Person{ "Jane", "Doe" }
	f.Println(v.GetFullName())
	f.Println(GetInitials(v))

	floatVar := MyFloat(-math.Sqrt2)
	f.Println(floatVar.Abs())

	intVar := MyInt(-234)
	f.Println(intVar.Abs())
}

