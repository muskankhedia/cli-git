package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadJsonFileData() RepoDetails {

	var res RepoDetails

	content, err := ioutil.ReadFile("store.json")
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal([]byte(content), &res)

	return res
}

func WriteJSONFileData(data RepoDetails) error {
	// now Marshal it
	result, error := json.Marshal(data)
	if error != nil {
		return error
	}
	err := ioutil.WriteFile("./store.json", result, 0644)
	if err != nil {
		return err
	}
	return nil
}
