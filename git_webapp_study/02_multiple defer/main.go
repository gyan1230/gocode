package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // print 4 3 2 1 0
	}
}
