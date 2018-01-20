package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

func main() {

	data, error := ioutil.ReadFile("/home/wangzhenya/3.csv")
	check(error)
	fileContent := string(data)

	lineArray := strings.Split(fileContent, "\n")

	for _,x:= range lineArray {
		if diff(x) {
			fmt.Println(x)
		}
	}

}

func diff(line string) bool  {
	if len(line) < 4 {
		fmt.Print(line)
		return false;
	}
	array:= strings.Split(line, ",")

	if array[0] != array[len(array) -1] {
		return true;
	}

	if array[1] != array[len(array) - 2] {
		return true;
	}


	return false;
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}


