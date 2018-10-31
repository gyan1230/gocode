package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:password@tcp(mydbinstance.ckzomjrwhtr8.us-east-2.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/details/", details)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "At index.")
	log.Println("Index passed...")
	check(err)
}

func details(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT idemployee FROM employee;`)
	check(err)
	defer rows.Close()

	var s, name string
	s = "records...\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		log.Println("error in opening DB...")
	}
}
