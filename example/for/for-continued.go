package _for

import "fmt"

func main()  {

	sum := 1
	for ;sum< 1000 ;  {
		sum += sum
	}

	fmt.Println(sum)
}
