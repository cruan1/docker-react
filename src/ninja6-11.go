package main

import (
	"fmt"
)

func main() {

	fmt.Println(factorial(5))

}

func factorial(i int) int {
	total := i

	if i == 0 {
		return 1
	}

	return total * factorial(i-1)
}

