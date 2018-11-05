package main

import (
	"fmt"
	"log"
)

func main() {
	n := 10
	var result int
	var ar [10]int
	ar[0] = 1
	ar[1] = 1
	ar[2] = 1
	ar[3] = 2
	ar[4] = 2
	ar[5] = 2
	ar[6] = 2
	ar[7] = 3
	ar[8] = 3
	ar[9] = 3

	result = pair(n, ar)
	fmt.Println("result is ", result)
}

func pair(n int, arr [10]int) int {
	var count, final, temp int

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				count++
				i = j
				log.Println("count is :", count)
				log.Println("i: ", i, "j: ", j)
			}
			temp = count

		}
		if temp > 1 {
			for {
				temp /= 2
				if temp == 0 {
					break
				}
				log.Println("temp is :", temp)
				final++
				log.Println("final is :", final)

			}
			//final += count % 2

		}

	}
	return final
}
