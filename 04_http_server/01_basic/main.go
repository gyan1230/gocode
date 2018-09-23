package main

import (
	"fmt"
	"net/http"
)

type gyan int

func (m gyan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello gyan!!!!!!")
}

func main() {
	var d gyan
	http.ListenAndServe(":8080", d)
}
