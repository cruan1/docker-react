package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	fmt.Println("I'm from main")

	go f1()

	go f2()

	wg.Wait()
}

func f1() {
	fmt.Println("I'm f1")
	wg.Done()
}

func f2() {
	fmt.Println("I want to tax all of you to death")
	wg.Done()
}

