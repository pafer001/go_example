package main

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func main()  {
	v := Vertex2{1, 3}
	p := &v
	fmt.Println(p)
	p.X = 1e9
	fmt.Println(v)
}
