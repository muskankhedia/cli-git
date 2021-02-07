package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
)

type repoDetail struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	PR      string `json:"pr"`
	Issues  string `json:"issues"`
	Visited int `json:"visited"`
}

type repoDetails struct {
	Details []repoDetail `json:"repos"`
}

func main() {

	var res repoDetails

	// Displays the project name text
	projectText()

	content, err := ioutil.ReadFile("store.json")
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal([]byte(content), &res)

	for {

		sort.Slice(res.Details[:], func(i, j int) bool {
			return res.Details[i].Visited > res.Details[j].Visited
		})

		funcMap := promptui.FuncMap
		funcMap["truncate"] = func(size int, input string) string {
			if len(input) <= size {
				return input
			}
			return input[:size-3] + "..."
		}

		funcMap["openBrowser"] = func(input string) bool {
			for x, v := range res.Details {
				if v.URL == input {
					res.Details[x].Visited += 1
				}
			}
			_ = browser.OpenURL(input)
			return true
		}

		templates := promptui.SelectTemplates{
			Active:   `🌁 {{ .Name | cyan | bold }}`,
			Inactive: `   {{ .Name | cyan }}`,
			Selected: `{{ "✔" | green | bold }} {{ "Repo Visited" | bold }}: {{ .Name | cyan }} {{ .URL | openBrowser }}`,
			Details: `Repo URL:
	{{ .URL | truncate 80 }}`,
		}

		list := promptui.Select{
			Label:             "Repo",
			Items:             res.Details,
			Templates:         &templates,
			StartInSearchMode: true,
			Searcher: func(input string, idx int) bool {
				repo := res.Details[idx]
				title := strings.ToLower(repo.Name)
				if strings.Contains(title, input) {
					return true
				}
				return false
			},
		}

		_, _, e := list.Run()
		if e != nil {
			os.Exit(1)
		}

	}

}
