package main

import "fmt"

func main() {
	type deemsucks int
	var x deemsucks
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
