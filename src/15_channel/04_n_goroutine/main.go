package main

import "fmt"

func main() {
	n := 3
	ch := make(chan int)
	sema := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 5; i++ {
				ch <- i
			}
			sema <- true
		}()

	}

	go func() {
		for i := 0; i < n; i++ {
			<-sema
		}
		close(ch)

	}()

	for n := range ch {
		fmt.Println(n)
	}
}
