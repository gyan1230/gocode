package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Details for db
type Details struct {
	ID   int
	Name string
}
type actor struct {
	ActorID    int
	FirstName  string
	LastName   string
	LastUpdate string
}

func main() {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Succesfully db connected...")

	// insert, err := db.Query("INSERT INTO `test`.`details` (`id`) VALUES ('5')")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// insert, err = db.Query("INSERT INTO `test`.`details` (`name`) VALUES ('ram')")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer insert.Close()
	// fmt.Println("Succesfully db Inserted...")

	result, err := db.Query("SELECT id FROM details")

	if err != nil {
		log.Fatalln(err)
	}

	result2, err := db.Query("SELECT name FROM details")
	if err != nil {
		log.Fatalln(err)
	}

	for result.Next() {
		var detail Details

		err = result.Scan(&detail.ID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(detail.ID)
	}
	for result2.Next() {
		var detail Details
		err = result2.Scan(&detail.Name)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(detail.Name)
	}

}
