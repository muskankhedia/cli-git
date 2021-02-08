package cmd

import (
	"github.com/muskankhedia/cli-git/pkg/common"
	"github.com/muskankhedia/cli-git/pkg/utils"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a particular repo in the list",
	Long: ` This command helps you to search among the list of added repos sorted
	based on recent search. You can use arrow keys to navigate in the list and enter 
	key to open the repo link in a new browser tab.`,
	Run: func(cmd *cobra.Command, args []string) {

		res := utils.LoadJsonFileData()
		for {
			common.SearchRepo(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
