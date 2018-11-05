package main

import (
	"io"
	"net/http"
)

func gyan(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi gyan...")
}

func chand(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello chand")
}

func (t something) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HI MOHANTY...")
}

type something int

func main() {

	var title something

	http.HandleFunc("/gyan/", gyan)
	http.HandleFunc("/chand", chand)
	http.Handle("/onlyhandle/", title)
	http.ListenAndServe(":8080", nil)
}
