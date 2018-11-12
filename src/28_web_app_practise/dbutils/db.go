package dbutils

import (
	"28_web_app_practise/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //
)

var (
	// Db global db connection
	Db *sql.DB
)

func init() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/cricbuzz")

	if err != nil {
		log.Fatalf("error in opening... %v", err)

	}
	Db = db
	// defer db.Close()

	err = Db.Ping()
	if err != nil {
		log.Println("error in ping...")
	}
	fmt.Println("DB connected...")

	return
}

//GetVideos :
func GetVideos() ([]*models.Details, error) {
	result, err := Db.Query("SELECT * FROM videos")

	if err != nil {
		log.Println("Error gttting videos from DB: ", err)
		return nil, err
	}
	defer result.Close()
	vid := []*models.Details{}
	for result.Next() {
		// detail := models.Details{}
		var id int
		var title, pub string
		err = result.Scan(&id, &title, &pub)
		if err != nil {
			log.Println("Error scanning videos from DB: ", err)
			return nil, err
		}
		detail := models.NewDetails(id, title, pub)
		vid = append(vid, detail)
	}
	return vid, nil
}
