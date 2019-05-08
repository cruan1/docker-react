package main

import (
	"fmt"
)

func main() {

	value := []int{1, 2, 3, 4, 5}
	value2 := []int{1, 2, 3, 4, 5, 6}

	defer fmt.Println(sum(value2...))

	fmt.Println(sum2(value))

}

func sum(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}
	return total
}

func sum2(x []int) int {
	total := 0

	for _, v := range x {
		total += v
	}
	return total
}
