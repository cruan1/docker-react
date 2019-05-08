package main

import (
	"fmt"
)

func main() {

	q := make(chan int)
	c := put(q)

	get(c)

}

func put(ch chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func get(d <-chan int) {
	for v := range d {
		fmt.Println(v)
	}
}
