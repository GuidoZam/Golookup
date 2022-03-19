package main

import (
	"fmt"

	"github.com/guidozam/utility"
	p "github.com/guidozam/utility/person"
)

func main() {
	v := p.Person { FirstName: "Jane", LastName: "Doe" }
	fmt.Println(v)
	utility.HelloWorld()
}