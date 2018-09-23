package main

import (
	"html/template"
	"io"
	"log"
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
func about(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "It's root")
}
func me(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.html")
	if err != nil {
		log.Fatalln("error in parsing", err)
	}
	err = tpl.ExecuteTemplate(res, "something.html", "gyan")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
