package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s\n\n", time.Now())
	io.WriteString(w, "Hello Go!")
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
