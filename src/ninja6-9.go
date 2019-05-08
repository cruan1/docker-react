package main

import (
	"fmt"
)

func main() {
	x := []int{1, 5, 10, 25, 30, 40, 41, 42}
	fmt.Println(sum(x...))

	fmt.Println(tens(sum, x...))
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func tens(f func(xi ...int) int, vi ...int) int {

	var yi []int

	for _, v := range vi {

		if v%10 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
