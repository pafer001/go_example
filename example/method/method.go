package main

import "fmt"

type TwoInt struct {
	a int
	b int
}

func main() {

	two1 := &TwoInt{12, 15}
	fmt.Println(two1.AddThem())
	fmt.Println(two1.AddToParams(20))

	two2 := new(TwoInt)
	two2.a = 1
	two2.b = 3

	fmt.Println(two2.AddThem())

	}

func (tn *TwoInt) AddThem() int  {

	return tn.a + tn.b
}

func (tn *TwoInt) AddToParams(param int ) int  {
	return tn.a + tn.b + param;
}



