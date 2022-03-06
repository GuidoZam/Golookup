package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	var x, y int = 3, 4
	var d = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(d)
	fmt.Println(x, y, d, z)
}