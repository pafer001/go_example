package main

import "fmt"

func main() {

	p := []int{2,3,5,7,11,14}
	fmt.Println("p ==", p)

	i := 0
	for i < len(p) {

		fmt.Printf("p[%d] == %d\n", i, p[i])
		i++
	}
}