package main

import (
	"fmt"
)

func main() {
	fdeem := struct {
		money    string
		location string
		CIO      bool
		weather  map[string]bool
		friends  []string
	}{
		money:    "ok",
		location: "sucks",
		CIO:      false,
		weather: map[string]bool{
			"sunny": false,
			"rain":  true,
			"snow":  false,
		},
		friends: []string{
			"Sanjeev",
			"Sudhakar",
			"Bret",
		},
	}
	fmt.Println(fdeem)
}
