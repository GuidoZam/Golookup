package main

import (
	f "fmt"
	m "math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func concatStrings(stringAFunc func() string, stringBFunc func() string) string {
	result := f.Sprint(stringAFunc(), stringBFunc())
	return result
}

func main() {
	hypot := func(x, y float64) float64 {
		return m.Sqrt(x*x + y*y)
	}
	f.Println(hypot(5, 12))

	f.Println(compute(hypot))
	f.Println(compute(m.Pow))

	concatenated := concatStrings(func() string { return "Hello"}, func() string { return "World"})
	f.Println(concatenated)
}