package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(vowels("Hi"))
}

func vowels(s string) int {
	re := regexp.MustComplie("a,e,i,o,u")
	fmt.Println(re.FindStringSubmatchIndex(s))
	return len(re.FindStringSubmatchIndex(s))
}
