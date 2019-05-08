package main

import "fmt"

type deemsucks int
var b deemsucks
var a int

func main() {
	a = 45
	fmt.Println(a)
        fmt.Printf("%T\n", a)
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int (b)
        fmt.Println(a)
        fmt.Printf("%T\n", a)
}
