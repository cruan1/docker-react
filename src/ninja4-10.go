package main

import "fmt"

func main() {

	m := map[string][]string{

		"Ruan_Alexander": []string{"Cars", "Legos", "Paw Patrol"},
		"Ruan_Kaitlyn":   []string{"Hatchimals", "Ponies", "Dolls"},
		"Ruan_Chris":     []string{"Games", "Computers", "Programs"},
	}
	fmt.Println(m)

	m["Xu_Fei"] = []string{"Shoes", "Make up", "Bags"}

	delete(m, "Xu_Fei")

	for i, v := range m {
		fmt.Println(i)
		for ii, vv := range v {
			fmt.Println(ii, vv)
		}
	}
}
