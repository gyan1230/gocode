package main

import "fmt"

var x int = 12

func main(){
	fmt.Println("Package level scope of a variable...")
	fmt.Println("Value of X is",x) // x can access here
	foo()
}

func foo(){
	fmt.Println("Value of X is",x) // x can acess here
}