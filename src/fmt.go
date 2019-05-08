package main

import "fmt"

var y = 4124

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#U\n", y)
	fmt.Printf("%#U\n%b\n%x\n", y, y, y)
        
	s := fmt.Sprintf("%#U\n%b\n%x", y, y, y)
	fmt.Println(s)

}
