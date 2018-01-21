package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

//map 映射键到值。
//map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
var m1 map[string]vertex

func main() {
	m1 = make(map[string]vertex)
	m1["Bell Labs"] = vertex{40.68433,-7439967}

	fmt.Println(m1["Bell Labs"])
}
