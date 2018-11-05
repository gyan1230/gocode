package main

import (
	"io"
	"net/http"
)

type gyan int

func main() {
	var root gyan
	//http.HandleFunc("/", root1)		//elegant way
	//http.Handle("/", http.HandlerFunc(root1)) // ok kind
	http.Handle("/index/", root) //for this ServeHTTP function is required implemented by user defined type
	http.ListenAndServe(":8080", nil)
}

func (root gyan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello GO....")
}

func root1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello root1 elegant way....")
}
