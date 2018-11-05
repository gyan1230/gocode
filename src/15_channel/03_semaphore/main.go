package main

import "fmt"

func main() {
	ch := make(chan int)
	sema := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		sema <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		sema <- true
	}()

	go func() {
		<-sema
		<-sema
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
