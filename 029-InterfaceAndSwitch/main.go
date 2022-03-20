package main

import (
	"fmt"
)

type Cat struct {
	FurrColor string
}

func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I'm a string! My value is '%s'\n", i.(string))
    case int:
        fmt.Printf("I'm an int! My value is %d\n", i.(int))
	case Cat:
		fmt.Printf("I'm a %s cat!\n", (i.(Cat)).FurrColor)
    default:
        fmt.Printf("Unknown type :(\n")
    }
}
func main() {  
    findType("This is a string!")
    findType(42)
    findType(3.14)
	findType(Cat{ FurrColor: "red" })
}