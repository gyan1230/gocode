package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type person struct {
	Firstname string
	Lastname  string
	Other     []string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/marshal/", mshl)
	http.HandleFunc("/encode", encd)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi i am in root")
}

func mshl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"Gyan",
		"Chand",
		[]string{"Slice", "of", "Strings"},
	}

	mshl, _ := json.Marshal(p1)
	w.Write(mshl)

}

func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"Gyan",
		"Chand",
		[]string{"I", "am", "Bad..."},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}

}
