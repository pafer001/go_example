package _switch

import (
	"fmt"
	"runtime"
)

func main()  {

	fmt.Println("go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os x")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("%s.", os)

	}
}
