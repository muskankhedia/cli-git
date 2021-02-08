package common

import (
	"errors"
	"net/url"
)

func ValidateEmptyInput(input string) error {
	_, err := url.ParseRequestURI(input)
	if err != nil && input != "exit" {
		return errors.New("Please input a valid URL")
	}
	return nil
}
