package main

import (
	"fmt"
)

func main() {

	s := dsucks()

	fmt.Println(s())

}

func dsucks() func() string {
	return func() string {
		return "d sucks"
	}
}
