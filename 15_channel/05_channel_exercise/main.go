package main

import (
	"fmt"
)

func printS(s string, n int) {
	fmt.Println(s, n)
}

func func1(ch, ch1, ch2 chan int) {
	for i := 0; i < 10; i++ {
		<-ch
		printS("go1", i)
		ch1 <- i
	}
	ch2 <- 0
}

func func2(ch, ch1 chan int) {
	for {
		i := <-ch1
		printS("go2", i)
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func1(ch, ch1, ch2)
	go func2(ch, ch1)
	ch <- 100
}
