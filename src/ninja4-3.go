package main

import "fmt"

func main() {

	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range a {

	fmt.Println(i, v)

}
        fmt.Printf("The type is %T\n", a)

	a1 := a[:5]
	fmt.Println(a1)

	a2 := a[5:]
	fmt.Println(a2)

	a3 := a[2:7]
        fmt.Println(a3)

        a4 := a[1:6]
        fmt.Println(a4)

}
