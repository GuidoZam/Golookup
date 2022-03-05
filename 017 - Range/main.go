package main

import "fmt"

var source = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// When using range there are two values returned: index and value of the current element
	for index, value := range source {
		fmt.Printf("source[%d] = %d\n", index, value)
	}

	pow := make([]int, 10)
	// You can avoid the declaration of the value
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// You can avoid the declaration of the index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}