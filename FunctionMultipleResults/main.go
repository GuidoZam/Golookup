package main

import "fmt"

func tryMeForSingleResult(x, y int) int {
	return x * y
}

func tryMeForMultipleResults(x, y int) (int, int) {
	return x * x, y * y 
}

func main() {
	singleResult := tryMeForSingleResult(2,6)
	fmt.Println("singleResult:\t\t", singleResult)

	xPow, yPow := tryMeForMultipleResults(3,7)
	fmt.Println("multipleResults x:\t", xPow)
	fmt.Println("multipleResults y:\t", yPow)
}