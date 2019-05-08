package main

import (
	"fmt"
)

func main() {

	y := func() string {
		return "hello"
	}()

	fmt.Println(y)

}
