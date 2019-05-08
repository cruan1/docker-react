package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {

		fmt.Printf("Number is %v, modulus is %v\n", i, i%4)
	}
}
