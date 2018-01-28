package main

/*
当一个函数在其函数体内调用自身，则称之为递归
 */

func NumFactorial(a int)(b int)  {

	if a <= 0 {
		 b = 1;
		 return
	}
	b = NumFactorial(a-1) * a
	return
}

func main() {
	print(NumFactorial(3))
}