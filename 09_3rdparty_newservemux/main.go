package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am in Root...")
}

func gyan(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi gyan...")
}

func main() {
	newRouter := httprouter.New
	newRouter.Get("/", index)
	newRouter.Get("/gyan", gyan)

	http.ListenAndServe(":8080", newRouter)
}
