package utils

type repoDetail struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	PR      string `json:"pr"`
	Issues  string `json:"issues"`
	Visited int    `json:"visited"`
}

type repoDetails struct {
	Details []repoDetail `json:"repos"`
}
