package main

import (
	"encoding/json"
	"fmt"
)

type thumbnail struct {
	GyanURL       string `json:"URL"`
	Height, Width int
}

type image struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	GyanIDs       []int `json:"IDs"`
}

func main() {
	rcv := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	var data image
	err := json.Unmarshal([]byte(rcv), &data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	for i, v := range data.GyanIDs {
		fmt.Println("ID[", i, "]", v)
	}
}
