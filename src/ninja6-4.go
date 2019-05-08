package main

import (
	"fmt"
)

func main() {

	p1 := person{
		first: "Chris",
		last:  "Ruan",
		age:   40,
	}
	
	p1.speak()

}

type person struct {
	first string
	last  string
	age   int
}

func (s person) speak() {
	fmt.Println("I am", s.first, s.last, ".  My age is", s.age)
}
