package main

import (
	"encoding/json"
	"flag"
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

func returnModifieds(changeds []string) string {
	var result string
	for index, value := range changeds {
		result += value
		if index+1 < len(changeds) {
			result += ","
		}
	}
	return result
}

func main() {
	jsonFile, _ := os.Open("/codefresh/volume/event.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var changes FilesChangeds
	json.Unmarshal(byteValue, &changes)
	result := changes.HeadCommit.Modified
	compare := flag.String("compare", "", "define operation mode")
	flag.Parse()
	if *compare != "" {
		fmt.Println(compareFiles(result, *compare))
	} else {
		fmt.Println(returnModifieds(result))
	}
}
