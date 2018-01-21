package main

import "fmt"

type Vertex4 struct {
	Lat, Long float64
}

var m  = map[string]Vertex4 {
	"Bell labs" : Vertex4{
		40.68433, -74.39967,
	},
	"Google" : Vertex4{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}