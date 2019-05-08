package main

import "fmt"

func main() {

	a := []string{"James", "Bond", "shaken, but not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	x := [][]string{a, b}

	fmt.Println(x)

	for _, v := range x {

		for ii, vv := range v {
			fmt.Println(ii, vv)
		}
	}

}
