package cmd

import (
	"fmt"

	"github.com/muskankhedia/cli-git/pkg/utils"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Check the customised repo list",
	Long: `Check the customised repo list`,
	Run: func(cmd *cobra.Command, args []string) {
		showAllRepos()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}

func showAllRepos() {

	var res utils.RepoDetails

	res = utils.LoadJsonFileData()

	for _, v := range res.Details {
		fmt.Println(v.URL)
	}

}
