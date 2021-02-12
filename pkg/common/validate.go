package common

import (
	"errors"
	"net/url"
	"strings"
)

func ValidateURLInput(input string) error {
	_, err := url.ParseRequestURI(input)
	if err != nil && input != "exit" {
		return errors.New("Please input a valid URL")
	}
	return nil
}

func ValidateEmptyInput(input string) error {
	if len(strings.TrimSpace(input)) < 1 {
		return errors.New("This input must not be empty")
	}
	return nil
}
