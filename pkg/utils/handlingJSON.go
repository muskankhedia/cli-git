package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const filePath = "store.json"

func fileExists(path string) bool {
	_, err := os.Open(path) // For read access.
	return err == nil
}

func LoadJsonFileData() RepoDetails {

	var res RepoDetails

	if !fileExists(filePath) {
		_, e := os.Create(filePath)
		if e != nil {
			log.Fatal(e)
		}
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal([]byte(content), &res)

	return res
}

func WriteJSONFileData(data RepoDetails) error {
	// now Marshal it
	result, error := json.MarshalIndent(data, "", "    ")
	if error != nil {
		return error
	}

	if !fileExists(filePath) {
		_, e := os.Create(filePath)
		if e != nil {
			log.Fatal(e)
		}
	}

	err := ioutil.WriteFile("./store.json", result, 0644)
	if err != nil {
		return err
	}
	return nil
}
