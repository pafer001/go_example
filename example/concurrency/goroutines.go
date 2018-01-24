package main

import (
	"time"
	"fmt"
)
/*
goroutine 是由 Go 运行时环境管理的轻量级线程。

go f(x, y, z)
开启一个新的 goroutine 执行

f(x, y, z)
f，x，y 和 z 是当前 goroutine 中定义的，但是在新的 goroutine 中运行 f。


 */
func say(s string)  {

	for i :=0; i<5;i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("word")
	say("hello")
}
