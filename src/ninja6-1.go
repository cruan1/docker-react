package main

import (
	"fmt"
)

func main() {

	fmt.Println(foo(5))
	
	fmt.Println(bar())
}

func foo(x int) int {
	x *= x
	return x
}

func bar() (int, string) {
	return 1, "This is stupid"
}
