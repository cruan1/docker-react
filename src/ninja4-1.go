package main

import "fmt"

func main() {

	a := [5]int{10, 12, 14, 16, 18}

	for i, v := range a {

	fmt.Println(i, v)

}
        fmt.Printf("The type is %T\n", a)

}
