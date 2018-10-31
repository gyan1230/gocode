package main

import "fmt"

func main() {
	var count int
	for count = 0; count < 10; count++ {
		if count == 5 {
			continue // print 1 to 9 except 5
			//break - print 1 2 3 4
		}
		fmt.Println(count)
	}

}
