package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 20)
	for i := 0; i < 50; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Length: ", len(mySlice), "capacity: ", cap(mySlice), "Value: ", mySlice[i])
	}
}
