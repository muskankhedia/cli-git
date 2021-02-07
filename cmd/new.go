package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/muskankhedia/cli-git/pkg/common"
	// "github.com/muskankhedia/cli-git/pkg/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

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

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add new URLs to the database",
	Long: `Add new URLs to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			url, err := common.PromptString("Repo URL")
			if(url == "exit"){
				break
			}
			if err != nil {
				panic(err)
			}
			addNewURL(url)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func addNewURL(url string) {
	fmt.Println("Added new URL")
	var res repoDetails

	res = LoadJsonFileData()

	ss := strings.Split(url, "/")
	s := ss[len(ss)-1]
	pulls := url + "/pulls"
	issues := url + "/issues"
	res.Details = append(res.Details, repoDetail{Name: s, URL: url, PR: pulls, Issues: issues, Visited: 0})

	// now Marshal it
	result, error := json.Marshal(res)
	if error != nil {
		panic(error)
	}
	err := ioutil.WriteFile("./store.json", result, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadJsonFileData() repoDetails {

	var res repoDetails

	content, err := ioutil.ReadFile("./store.json")
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal([]byte(content), &res)

	return res
}