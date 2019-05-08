package main

import (
	"fmt"
)

func main() {
	enclose()
}

func enclose() {
	x := 1
	fmt.Println("x is now", x)

	{
		x := 2
		fmt.Println("x is now", x)
	}

	fmt.Println("x is now", x)
}

