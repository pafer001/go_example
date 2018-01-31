package main

import "fmt"

type Shaper1 interface {
	Area() float32
}

type Square1 struct {
	side float32
}

func (sq * Square1)Area()float32  {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle)Area() float32  {
	return r.length * r.length
}

func main() {

	r := Rectangle{5, 3}
	q := &Square1{5}

	shapers := []Shaper1{r, q}
	for n, _ :=range  shapers{

		fmt.Println(shapers[n])
		fmt.Println(shapers[n].Area())
	}
}
