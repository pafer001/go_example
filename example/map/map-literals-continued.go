package _map

import "fmt"

type Vertex5 struct {
	Lat, Long float64
}

var m  = map[string] Vertex5 {
	"Bell labs" :{40,67},
	"Google" : {45, 89},
}

func main() {

	fmt.Println(m)
}
