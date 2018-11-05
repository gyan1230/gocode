package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	child := []int{11, 12, 13, 14, 15}
	var total int
	intSlice = append(intSlice, 10)
	intSlice = append(intSlice, child...)
	for _, val := range intSlice {
		total += val
		fmt.Println(val)
	}

	fmt.Println("Sum of value is ", total)
	fmt.Println("Avg is :", total/len(intSlice))
}
