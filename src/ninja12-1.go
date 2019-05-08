package main

import (
	"fmt"
	"dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Cubby",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
