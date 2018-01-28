package main

import (
	"reflect"
	"fmt"
)

type TagType struct {

	field1 bool "An important answer"
	field2 string "The name of the thing"
	field3 int "How much there are"
}

func main() {

	tagType := TagType{true, "Barak obama", 1}

	for i:=0; i<3; i++ {
		refTag(tagType, i)
	}
}

func refTag(tt TagType, ix int )  {

	ttType := reflect.TypeOf(tt)
	field := ttType.Field(ix)
	fmt.Printf("%v\n", field.Tag)
}
