package cmd

import (
	"fmt"

	"github.com/muskankhedia/cli-git/pkg/common"
	"github.com/muskankhedia/cli-git/pkg/utils"
	"github.com/spf13/cobra"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new URLs to the database",
	Long:  `Add new URLs to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			url, err := common.PromptAddURL("Repo URL")
			if url == "exit" {
				break
			}
			if err != nil {
				break
			}
			addNewURL(url)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addNewURL(url string) {
	var res utils.RepoDetails

	res = utils.LoadJsonFileData()

	for _, v := range res.Details {
		if(v.URL == url) {
			fmt.Println("This URL already exists, try adding another URL")
			return
		}
	}

	repoURLArray := strings.Split(url, "/")
	repoName := repoURLArray[len(repoURLArray)-1]
	pulls := url + "/pulls"
	issues := url + "/issues"
	res.Details = append(res.Details, utils.RepoDetail{Name: repoName, URL: url, PR: pulls, Issues: issues, Visited: 0})

	err := utils.WriteJSONFileData(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Added new URL")
}
