package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", index)
	http.Handle("/", http.HandlerFunc(index))
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", gyan)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, I am in root index...")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, I am in dog route..")
}

func gyan(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "something.gohtml", "gyan")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
}
