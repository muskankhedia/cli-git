package common

import (
	"github.com/manifoldco/promptui"
)

func PromptURL(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: ValidateEmptyInput,
	}

	return prompt.Run()
}
