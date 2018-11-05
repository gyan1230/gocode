package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mutex sync.Mutex
var count int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Final count value is: ", count)
}

func foo() {
	for i := 0; i < 100; i++ {
		//	mutex.Lock()
		//count++
		atomic.AddInt64(&count, 1)
		fmt.Println("foo :", atomic.LoadInt64(&count))
		//	mutex.Unlock()

	}

	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		//mutex.Lock()
		//count++
		atomic.AddInt64(&count, 1)
		fmt.Println("Bar :", atomic.LoadInt64(&count))
		//mutex.Unlock()

	}
	wg.Done()
}
