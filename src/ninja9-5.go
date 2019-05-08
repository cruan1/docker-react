package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var x int64
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
		atomic.AddInt64(&x, 1)
		runtime.Gosched()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())

	}
	wg.Done()
}

func increment2() {
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&x, 1)
		runtime.Gosched()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())

	}
	wg.Done()
}
