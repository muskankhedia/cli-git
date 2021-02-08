package utils

type RepoDetail struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	PR      string `json:"pr"`
	Issues  string `json:"issues"`
	Visited int    `json:"visited"`
}

type RepoDetails struct {
	Details []RepoDetail `json:"repos"`
}
