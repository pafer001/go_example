package main

import "fmt"

type struct1 struct {
	i1 int
	f1 float32
	str string
}

func main() {

	ms := new(struct1)
	ms.f1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("the int is:%d\n", ms.i1)
	fmt.Printf("the float is:%f\n", ms.f1)
	fmt.Printf("the strings is:%s\n", ms.str)

	fmt.Println(ms)
}
