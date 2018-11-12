package main

import (
	"28_web_app_practise/dbutils"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func getvideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	videos, err := dbutils.GetVideos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(videos)
	if err != nil {
		log.Println("Error marshaling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
	return

}

func main() {
	http.HandleFunc("/videos/", getvideos)
	http.ListenAndServe(":8080", nil)
}
