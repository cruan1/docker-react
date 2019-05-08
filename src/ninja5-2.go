package main

import (
	"fmt"
)

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {

	p1 := person{
		first:    "Chris",
		last:     "Ruan",
		iceCream: []string{"Toasted Almond", "Vanilla", "Mango", "Red Bean"},
	}

	p2 := person{
		first:    "Fei",
		last:     "Xu",
		iceCream: []string{"Vanilla", "Chocolate"},
	}

	fmt.Println(p1.first, p1.last)

	for _, v := range p1.iceCream {

		fmt.Println(v)

	}

	fmt.Println(p2.first, p2.last)

	for _, v := range p2.iceCream {

		fmt.Println(v)

	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, val := range v.iceCream{
		fmt.Println(val)
		}

	}
}
