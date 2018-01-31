package main

import "fmt"

type Person1 struct {

	firstName string
	lastName string
}

func (p *Person1)FirstName()string  {
	return p.firstName
}

func (p *Person1)SetFirstName(newString string)  {
	p.firstName = newString
}

func main() {
	p := new(Person1)

	p.SetFirstName("Eric")
	fmt.Println(p.FirstName())
}





