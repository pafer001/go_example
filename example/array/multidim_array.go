package main

import "fmt"

const (
	WIDTH = 1920
	HEIGHT = 1080
)

/*
数组通常是一维的，但是可以用来组装成多维数组，例如：[3][5]int，[2][2][2]float64。
内部数组总是长度相同的。Go 语言的多维数组是矩形式的
 */
var screen [WIDTH][HEIGHT] int

func main() {
	for y :=0;y<HEIGHT ;y++  {

		for x :=0;x<WIDTH ;x++  {
			screen[x][y] = x+y
		}
	}

	fmt.Println(screen)
}