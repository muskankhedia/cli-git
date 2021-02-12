package common

import (
	"github.com/manifoldco/promptui"
)

func PromptURL(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: ValidateURLInput,
	}

	return prompt.Run()
}

func GetUsernamePrompt() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Enter your Github Username",
		Validate: ValidateEmptyInput,
	}

	return prompt.Run()
}
