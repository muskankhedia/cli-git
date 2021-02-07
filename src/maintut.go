package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
)

type repoDetail struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	PR      string `json:"pr"`
	Issues  string `json:"issues"`
	Visited string `json:"visited"`
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

		funcMap := promptui.FuncMap
		funcMap["truncate"] = func(size int, input string) string {
			if len(input) <= size {
				return input
			}
			return input[:size-3] + "..."
		}

		funcMap["openBrowser"] = func(input string) bool {
			_ = browser.OpenURL(input)
			return true
		}

		templates := promptui.SelectTemplates{
			Active:   `🌁 {{ .Name | cyan | bold }}`,
			Inactive: `   {{ .Name | cyan }}`,
			Selected: `{{ "✔" | green | bold }} {{ "Repo" | bold }}: {{ .Name | cyan }} {{ .URL | openBrowser }}`,
			Details: `Repo URL:
	{{ .URL | truncate 80 }}`,
		}

		list := promptui.Select{
			Label:     "Repo",
			Items:     res.Details,
			Templates: &templates,
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

		idx, _, err := list.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(res.Details[idx].URL)
	}

}
