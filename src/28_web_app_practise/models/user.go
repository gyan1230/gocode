package models

//Details for DB
type Details struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Publish string `json:"publish"`
}

//NewDetails :
func NewDetails(id int, title, pub string) *Details {
	return &Details{ID: id, Title: title, Publish: pub}
}
