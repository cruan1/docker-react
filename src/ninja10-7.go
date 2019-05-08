package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	go func() {
		const goroutine = 10
		var wg sync.WaitGroup
		wg.Add(goroutine)

		for i := 0; i < goroutine; i++ {

			go func() {
				for i := 0; i < 10; i++ {
					c <- i
				}
				wg.Done()

			}()
		}

		wg.Wait()
		close(c)
	}()
	
	for v := range c {
		fmt.Println(v)
	}
}
