package main

import "fmt"

//变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的。

func f2() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	fmt.Println(f2())
}