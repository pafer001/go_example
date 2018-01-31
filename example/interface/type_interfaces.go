package main

import (
	"math"
	"fmt"
)

type Square2 struct {
	side float32
}

type Circle2 struct {
	radius float32
}

type Shaper2 interface {
	Area() float32
}

func main() {

	var areaIntf Shaper2
	sq1 := new(Square2)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*Square2); ok {
		fmt.Println("tepe is:%T\n", t)
	}


	if u, ok := areaIntf.(*Circle2); ok {
		fmt.Println("tepe is:%T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}


	switch t := areaIntf.(type) {
	case *Square2:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle2:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

}

func (sq *Square2)Area() float32{
	return sq.side * sq.side
}

func (ci * Circle2)Area() float32  {
	return ci.radius * ci.radius * math.Pi
}
