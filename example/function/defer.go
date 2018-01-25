package main

import "fmt"

/*

关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。

关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。

当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：


 */
func function1()  {
	fmt.Println("start in function1.....")
	defer function2()
	fmt.Println("end in function1.....")
}

func function2()  {
	fmt.Println("function2:defer")
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func main() {
	function1()
	f()
}
