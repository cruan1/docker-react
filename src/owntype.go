package main

import "fmt"

type deemsucks int
var b deemsucks

func main() {
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
