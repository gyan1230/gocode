package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}

	}()
	time.Sleep(1 * time.Second)
}
