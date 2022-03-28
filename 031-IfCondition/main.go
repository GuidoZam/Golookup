package main

import (
	"fmt"
)

func ifWithParenthesis(n int) {
	if (n > 3) {
		fmt.Println("Number is major than 3!")
	}
}

func ifWithoutParenthesis(n int) {
	if n % 2 == 1 {
		fmt.Println("Number is odd!")
	}
}

func ifWithShortStatement(n int) {
	if result := n * 2; result % 2 == 0 {
		fmt.Println("Result is even!")
	}
}

func ifWithShortStatementAndElse(n int) {
	if result := n * 2; result > 10 {
		fmt.Printf("Result is %v and is grater than 10!\n", result)
	} else {
		fmt.Printf("Result is %v and is lower than 10!\n", result)
	}
}

func main() {
	n := 5
	fmt.Printf("Current number is %v\n", n)

	ifWithParenthesis(n)
	ifWithoutParenthesis(n)
	ifWithShortStatement(n)
	ifWithShortStatementAndElse(2)
	ifWithShortStatementAndElse(9)
}