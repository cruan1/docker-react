package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("This is an error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "Chris",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}