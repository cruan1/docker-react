package main

import "fmt"

var y = 42
var z = `I said "Deem sucks"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
