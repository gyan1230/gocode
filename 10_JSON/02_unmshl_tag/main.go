package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string
	Lastname  string
	Other     []string
}

func main() {
	var p1 person
	rcv := `{"Firstname":"Gyan","Lastname":"Chand","Other":["slice","of","string"]}`
	json.Unmarshal([]byte(rcv), &p1)

	fmt.Println(p1)
	fmt.Printf(p1.Lastname)

}
