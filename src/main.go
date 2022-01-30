package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type FilesChangeds struct {
	HeadCommit struct {
		Modified []string `json:"modified"`
	} `json:"head_commit"`
}

func compareFiles(changeds []string, fileName string) {
	for _, value := range changeds {
		if value == fileName {
			os.Exit(0)
		}
	}

	os.Exit(1)
}

func main() {
	jsonFile, _ := os.Open("/codefresh/volume/event.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var teste FilesChangeds
	json.Unmarshal(byteValue, &teste)
	result := teste.HeadCommit.Modified
	compareFiles(result, os.Args[1])
}
