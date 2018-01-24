package main

import (
	"strings"
	"fmt"
	"io"
)

func main()  {

	reader := strings.NewReader("Hello, Reader")
	b := make([]byte, 16)
	for true {
		n, err := reader.Read(b)
		fmt.Printf("n= %v err =%v\n", n ,err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
