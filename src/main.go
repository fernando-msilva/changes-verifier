package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FilesChangeds struct {
	HeadCommit struct {
		Modified []string `json:"modified"`
	} `json:"head_commit"`
}

func compareFiles(changeds []string, fileName string) string {
	for _, value := range changeds {
		if value == fileName {
			return "true"
		}
	}

	return "false"
}

func main() {
	jsonFile, _ := os.Open("/codefresh/volume/event.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var teste FilesChangeds
	json.Unmarshal(byteValue, &teste)
	result := teste.HeadCommit.Modified
	fmt.Println(compareFiles(result, os.Args[1]))
}
