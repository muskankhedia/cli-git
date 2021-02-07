package main

import (
	"encoding/json"
	"errors"
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
		// Prompts the user to search for the repository
		prompt := promptui.Prompt{
			Label: "Search",
			Validate: func(input string) error {
				if len(input) < 3 {
					return errors.New("Search term must have at least 3 characters")
				}
				return nil
			},
		}

		// Runs the searched word
		keyword, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var selectedRepos []repoDetail

		for i := 0; i < len(res.Details); i++ {
			if strings.Contains(res.Details[i].Name, keyword) {
				selectedRepos = append(selectedRepos, res.Details[i])
			}
		}

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
			Items:     selectedRepos,
			Templates: &templates,
			Searcher: func(input string, idx int) bool {
				recipe := selectedRepos[idx]
				title := strings.ToLower(recipe.Name)
				fmt.Println("input: ", input)
				if strings.Contains(title, input) {
					return true
				}
				if strings.Contains(recipe.URL, input) {
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

		fmt.Println(selectedRepos[idx].URL)
	}

}
