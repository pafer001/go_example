package main

import "fmt"
//你可以取任意数组常量的地址来作为指向新实例的指针。
func fp1(a *[3]int) {
	fmt.Println(a)
}

func main() {
	for i := 0; i < 3; i++ {
		fp1(&[3]int{i, i * i, i * i * i})
	}
}
