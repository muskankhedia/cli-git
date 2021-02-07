package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

// Model
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func githubOAuth() {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "66692acdcd801ebd2a0327d99b0caa3e9ba336ba"},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)
	opt := &github.RepositoryListByOrgOptions{Type: "private"}
	repos, _, err := client.Repositories.ListByOrg(context, "interviewstreet", opt)
	// repo, _, err := client.Repositories.Get(context, "muskankhedia", "Jarvis-Personal-Assistant")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	// var packs []Package

	for _, v := range repos {
		// fmt.Printf("%+v\n", v)

		pack := &Package{
			FullName:    *v.FullName,
			Description: *v.Description,
			ForksCount:  *v.ForksCount,
			StarsCount:  *v.StargazersCount,
		}
		fmt.Printf("%+v\n", pack)

		// packs = append(packs, *pack)
	}

	// pack := &Package{
	// 	FullName: *repo.FullName,
	// 	Description: *repo.Description,
	// 	ForksCount: *repo.ForksCount,
	// 	StarsCount: *repo.StargazersCount,
	// }

	// fmt.Printf("%+v\n", packs)

	// fmt.Println("Problem: ", repos)
}
