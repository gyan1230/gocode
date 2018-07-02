package main

import "fmt"

func main(){

	a := 10
	b := "gyan"
	c := 12.30
	d := false
	var a1 int
	var b1 string
	var c1 float64
	var d1 bool


	fmt.Printf("%v : %T\n",a,a)
	fmt.Printf("%v : %T\n",b,b)
	fmt.Printf("%v : %T\n",c,c)
	fmt.Printf("%v : %T\n\n",d,d)
	fmt.Printf("default value of int :%v\n",a1,)
	fmt.Printf("default value of string :%v\n",b1,)
	fmt.Printf("default value of float :%v\n",c1,)
	fmt.Printf("default value of bool :%v\n",d1,)
}