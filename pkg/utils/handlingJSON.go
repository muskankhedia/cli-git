package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadJsonFileData() repoDetails {

	var res repoDetails

	content, err := ioutil.ReadFile("store.json")
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal([]byte(content), &res)

	return res
}