package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/", about)

	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}
func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "It's me gyan")
}
func about(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "It's root")
}
