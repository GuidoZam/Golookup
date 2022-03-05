package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var secondMap = map[string]Vertex{
	// You can specify the type
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	// Or you can avoid the explicit type
	"Google": { 37.42202, -122.08408 },
}

func main() {
	m = make(map[string]Vertex)
	// Add entry
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	// Add entry
	m["Meh"] = Vertex{
		12.3233, 54.10267,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Meh"])

	fmt.Println(secondMap)
}