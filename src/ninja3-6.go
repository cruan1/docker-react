package main

import "fmt"

func main() {
	i := 1
	if i == 20 {
		fmt.Println("i = 20")
	} else if i == 21 {
		fmt.Println("i = 21")
	} else {
		fmt.Println("i != 20,21")
	}
}
