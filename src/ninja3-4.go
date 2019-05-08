package main

import "fmt"

func main() {
	b := 1978
	for {
		fmt.Println(b)
		b++
		if b > 2019 {
			break
		}
	}

}
