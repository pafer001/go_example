package main

import "fmt"

func main() {

	slice1 := make([]int, 0, 10)

	for i:=0;i<cap(slice1) ;i++  {
		slice1 = slice1[0:i+1]
		slice1[i] = i
		fmt.Printf("the length of  slice is %d\n", len(slice1))
	}

	for i :=0; i< len(slice1) ;i++  {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}

	var ar = [10] int{1,2,4,5}
	fmt.Println(ar[len(ar):len(ar)])
}
