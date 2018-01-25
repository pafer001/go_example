package main

import "fmt"

func Add(a, b int, sum *int)  {
	*sum =  a + b
}

/*
传递指针给函数不但可以节省内存（因为没有复制变量的值）
，而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回
 */
func main() {
	n := 0
	sum := &n
	Add(1, 2, sum)

	n = *sum +17
	fmt.Println("sum is ", *sum +1)

}