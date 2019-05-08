package main

import (
	"fmt"
)

func main() {

	var p1 = person{
		first: "Chris",
		last:  "Ruan",
		age:   40,
	}

	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)

}

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).first = "Bob"
	(*p).last = "Dole"
	(*p).age = 99
}
