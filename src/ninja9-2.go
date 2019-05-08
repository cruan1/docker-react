package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hi, I'm", p.first, p.last, "\nMy age is", p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	fmt.Println("Hello, playground")

	p1 := person{
		first: "Chris",
		last:  "Ruan",
		age:   40,
	}

	p1.speak()
	
	saySomething(&p1)
}
