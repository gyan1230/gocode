package main

import (
	"19_cricbuzz/data"
	"encoding/json"
	"fmt"

	"net/http"
)

func main() {
	http.HandleFunc("/scores", score)
	http.HandleFunc("/matches", detail)

	http.ListenAndServe(":8080", nil)

}

func score(w http.ResponseWriter, r *http.Request) {
	m := data.MatchScore{
		BatScore:  "155/1",
		BallScore: "16.4 ovrs",
	}
	mj, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", mj)
	return
}

func detail(w http.ResponseWriter, r *http.Request) {
	d := data.Team{
		FirstTeam:  "India",
		SecondTeam: "Pakistan",
	}

	//fmt.Fprintf(w, "%s\n", dj)
	//w.Write(dj)
	err2 := json.NewEncoder(w).Encode(d)
	if err2 != nil {
		fmt.Println(err2)
	}
}
