package data

//Team is exported function
type Team struct {
	FirstTeam  string `json:"firstTeam"`
	SecondTeam string `json:"secondTeam"`
}

//MatchScore is exported function
type MatchScore struct {
	BatScore  string `json:"BatScore"`
	BallScore string `json:"BallScore"`
}
