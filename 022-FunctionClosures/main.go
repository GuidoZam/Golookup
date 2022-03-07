package main

import f "fmt"

func incrementor() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		f.Printf("adder() with value %v summed to %v\n", x, sum)
		return sum
	}
}

func main() {
	next := incrementor()
	f.Println(next())
	f.Println(next())
	f.Println(next()) 

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		f.Println(
			pos(i),
			neg(-2*i),
		)
	}
}