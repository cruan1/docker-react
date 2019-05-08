package main

import (
	"fmt"
	"runtime"
	"sync"
)

var x int
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go increment()
	go increment2()
	wg.Wait()
	fmt.Println(x)
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

}

func increment() {
	for i := 0; i < 100; i++ {
		y := x
		runtime.Gosched()
		y++
		x = y
		fmt.Println("Goroutines\t", runtime.NumGoroutine())

	}
	wg.Done()
}

func increment2() {
	for i := 0; i < 100; i++ {
		y := x
		runtime.Gosched()
		y++
		x = y
		fmt.Println("Goroutines\t", runtime.NumGoroutine())

	}
	wg.Done()
}
