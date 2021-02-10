package common

import (
	"os"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/muskankhedia/cli-git/pkg/utils"
	"github.com/pkg/browser"
)

func SearchRepo(res utils.RepoDetails) {
	sort.Slice(res.Details[:], func(i, j int) bool {
		return res.Details[i].Visited > res.Details[j].Visited
	})

	funcMap := promptui.FuncMap
	funcMap["truncate"] = func(size int, input string) string {
		if len(input) <= size {
			return input
		}
		return input[:size-3] + "..."
	}

	funcMap["openBrowser"] = func(input string) bool {
		for x, v := range res.Details {
			if v.URL == input {
				res.Details[x].Visited += 1
				_ = utils.WriteJSONFileData(res)
			}
		}
		_ = browser.OpenURL(input)
		return true
	}

	templates := promptui.SelectTemplates{
		Active:   `▶ {{ .Name | cyan | bold }}`,
		Inactive: `   {{ .Name | cyan }}`,
		Selected: `{{ "✔" | green | bold }} {{ "Repo Visited" | bold }}: {{ .Name | cyan }} {{ .URL | openBrowser }}`,
		Details: `Repo URL:
{{ .URL | truncate 80 }}`,
	}

	selectKeys := promptui.SelectKeys{
		Prev:     promptui.Key{Code: promptui.KeyPrev, Display: promptui.KeyPrevDisplay},
		Next:     promptui.Key{Code: promptui.KeyNext, Display: promptui.KeyNextDisplay},
		PageUp:   promptui.Key{Code: utils.KeyBackward, Display: utils.KeyBackwardDisplay},
		PageDown: promptui.Key{Code: utils.KeyForward, Display: utils.KeyForwardDisplay},
		Search:   promptui.Key{Code: '/', Display: "/"},
	}

	list := promptui.Select{
		Label:             "Repo",
		Items:             res.Details,
		Templates:         &templates,
		StartInSearchMode: true,
		Keys:              &selectKeys,
		Searcher: func(input string, idx int) bool {
			repo := res.Details[idx]
			title := strings.ToLower(repo.Name)
			return strings.Contains(title, input)
		},
	}

	_, _, e := list.Run()
	if e != nil {
		os.Exit(1)
	}
}
