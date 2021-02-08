package cmd

import (
	"fmt"

	"github.com/muskankhedia/cli-git/pkg/utils"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
