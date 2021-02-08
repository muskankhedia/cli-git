package cmd

import (
	"fmt"

	"github.com/muskankhedia/cli-git/pkg/common"
	"github.com/muskankhedia/cli-git/pkg/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repo from the list",
	Long:  `Delete a repo from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			url, err := common.PromptURL("Repo URL")
			if url == "exit" {
				break
			}
			if err != nil {
				break
			}
			deleteURL(url)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteURL(url string) {
	var res utils.RepoDetails

	res = utils.LoadJsonFileData()

	for i, v := range res.Details {
		if v.URL == url {
			res.Details = append(res.Details[:i], res.Details[i+1:]...)
			err := utils.WriteJSONFileData(res)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("The URL has been deleted")
			return
		}
	}

	fmt.Println("This URL doesn't exist")
}
