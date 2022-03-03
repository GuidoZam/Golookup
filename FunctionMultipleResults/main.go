package main

import "fmt"

func tryMeForSingleResult(x, y int) int {
	return x * y
}

func tryMeForMultipleResults(x, y int) (int, int, string) {
	return x * x, y * y, fmt.Sprintf("%v%v", x, y)
}

func main() {
	x := 3
	y := 7

	fmt.Println("param x:\t\t", x)
	fmt.Println("param y:\t\t", y)

	singleResult := tryMeForSingleResult(x,y)
	fmt.Println("singleResult:\t\t", singleResult)

	xPow, yPow, concatenated := tryMeForMultipleResults(x,y)
	fmt.Println("multipleResults x:\t", xPow)
	fmt.Println("multipleResults y:\t", yPow)
	fmt.Println("multipleResults concat:\t", concatenated)
}