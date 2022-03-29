package main

import "fmt"

type I interface {
	CustomMethod()
}

type T struct {
	StringValue string
}

// Interface implementation on type T
func (t *T) CustomMethod() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.StringValue)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.CustomMethod()

	i = &T{"hello"}
	describe(i)
	i.CustomMethod()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}