package main

import (
	"fmt"
)

func main() {
	const (
		//_ = iota
		a = 2019 - iota
		b = a - iota
		c = a - iota
		d = a - iota
	)
	fmt.Println(a, b, c, d)
}
