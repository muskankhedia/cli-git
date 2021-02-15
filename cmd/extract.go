package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/muskankhedia/cli-git/pkg/common"
	"github.com/muskankhedia/cli-git/pkg/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type userGithubRepo struct {
	URL string `json:"html_url"`
}

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract all the repos from the specified github username",
	Long: `Extract all the repos from the specified github username`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := common.GetUsernamePrompt()
		extractGithubData(username)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}

func extractGithubData(username string) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Color("red")
	s.Prefix = "Extracting Repo data from your profile: "

	s.Start()
	count := 0

	for {
		count++
		getUrl := "https://api.github.com/users/" + username + "/repos?per_page=100&page=" + strconv.Itoa(count)
		resp, err := http.Get(getUrl)
		if err != nil {
			fmt.Println("User Limit Exceeded")
			return
		}

		responseBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Printf("Could not read response body. %v", err)
		}

		var listGithubRepos []userGithubRepo

		if err := json.Unmarshal([]byte(responseBytes), &listGithubRepos); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		}

		if len(listGithubRepos) == 0 {
			break
		}

		addURLsToStore(listGithubRepos)

		fmt.Println("Added URLs from github repo")
	}

	s.Stop()
}

func addURLsToStore(listGithubRepos []userGithubRepo) {
	var res utils.RepoDetails
	res = utils.LoadJsonFileData()

	for _, url := range listGithubRepos {

		for _, v := range res.Details {
			if v.URL == url.URL {
				continue
			}
		}

		repoURLArray := strings.Split(url.URL, "/")
		repoName := repoURLArray[len(repoURLArray)-1]
		pulls := url.URL + "/pulls"
		issues := url.URL + "/issues"
		res.Details = append(res.Details, utils.RepoDetail{Name: repoName, URL: url.URL, PR: pulls, Issues: issues, Visited: 0})
	}
	err := utils.WriteJSONFileData(res)
	if err != nil {
		fmt.Println(err)
	}
}
