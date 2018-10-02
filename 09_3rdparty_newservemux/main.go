package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "I am in Root...")
}

func Gyan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Hi gyan...")
}

func main() {
	newRouter := httprouter.New()
	newRouter.GET("/", Index)
	newRouter.GET("/gyan", Gyan)

	http.ListenAndServe(":8080", newRouter)
}
