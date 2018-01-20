package main

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

var (
	v1 =Vertex3{1, 2}
	v2 = Vertex3{X:1}
	v3 = Vertex3{}
	p = &Vertex3{1, 2}
)

func main()  {

	fmt.Println(v1, p, v3, v3, p.X)
}

