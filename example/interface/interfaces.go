package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}
//接口
//接口类型是由一组方法定义的集合。
//
//接口类型的值可以存放实现这些方法的任何值。

func main() {

	var a Abser;
	f := MyFloat1(-math.Sqrt2)
	v := Vertex1{3,4}

	a=f
	a=&v

	a=v
	fmt.Println(a.Abs())
}

type MyFloat1 float64

func (f MyFloat1) abs() float64 {

	if f <0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex1 struct {
	X, Y float64
}
