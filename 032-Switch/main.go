package main

import (
	"fmt"
	"runtime"
	"time"
)

func simpleSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchEvaluationOrder(n int) {
	switch n {
		case 0:
			fmt.Println("Number is 0!")
		case 2:
			fmt.Println("Number is 2!")
	}
}

func switchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Guten Morgen!")
	case t.Hour() < 18:
		fmt.Println("Guten Tag!")
	case t.Hour() < 21:
		fmt.Println("Guten Abend!")
	default:
		fmt.Println("Gute Nacht!")
	}
}

func main() {
	simpleSwitch()
	switchEvaluationOrder(0)
	switchEvaluationOrder(2)
	switchWithNoCondition()
}