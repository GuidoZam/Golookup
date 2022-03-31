package main

import (
	"fmt"
)

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("An exception occurred: %s", e.Message)
}

func isGreater_Internal(a int, b int) (string, error) {
	if a > b {
		return fmt.Sprintf("%d is greater than %d", a, b), nil
	} else if a < b {
		return fmt.Sprintf("%d is lower than %d", a, b), nil
	} else {
		return "", &MyError{"Numbers cannot be equals"}
	}
}

func isGreater(a int, b int) {
	if result, err := isGreater_Internal(a, b); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
}

func main() {
	isGreater(3, 6)
	isGreater(42, 17)
	isGreater(8, 8)
}

