package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe bool = false
	MaxInt uint64
	z complex128 = cmplx.Sqrt(-5+12i)
)

func main()  {

	const f = "%T(%v)\n"
	fmt.Println(f, ToBe, ToBe)
	fmt.Println(f, MaxInt, MaxInt)
	fmt.Println(f, z, z)
}
