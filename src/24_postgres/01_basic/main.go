package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Book struct for postgres
type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://gyan:password@localhost/bookstore?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("You are connected...")

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}

		err = rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			fmt.Println(err)
		}
		bks = append(bks, bk)
	}
	for _, v := range bks {
		fmt.Printf("%s ,%s, %s, %.2f\n", v.isbn, v.title, v.author, v.price)
	}
}
