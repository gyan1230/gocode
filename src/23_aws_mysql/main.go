package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

//Info struct for json
type Info struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	db, err = sql.Open("mysql", "gyanaws:password@tcp(gyandbinstance.ckzomjrwhtr8.us-east-2.rds.amazonaws.com:3306)/test?charset=utf8")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/details/", detail)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	io.WriteString(w, "At index")

}

func detail(w http.ResponseWriter, r *http.Request) {
	info := []Info{}
	rows, err := db.Query(`SELECT name,id,age FROM test.details`)
	w.Header().Set("Content-Type", "application/json")
	if check(err) {
		log.Println("ERROR GETTING DATA FROM DB: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	var id, age int
	var name string

	for rows.Next() {
		err = rows.Scan(&name, &id, &age)
		if check(err) {
			log.Println("ERROR Scanning DATA FROM DB: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		info = append(info, Info{id, name, age})
	}
	d, err := json.Marshal(&info)
	if check(err) {
		log.Println("ERROR Marshal ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(d)
	return
}

func check(err error) bool {
	if err != nil {
		return true
	}
	return false
}
