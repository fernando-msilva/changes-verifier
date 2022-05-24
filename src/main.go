package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type FilesChangeds struct {
	HeadCommit struct {
		Modified []string `json:"modified"`
		Added    []string `json:"added"`
	} `json:"head_commit"`
}

func compare(changeds []string, arg string) string {
	result := "false"
	if arg[len(arg)-1:] == "/" {
		result = comparePaths(changeds, arg)
	} else {
		result = compareFiles(changeds, arg)
	}

	return result
}

func compareFiles(changeds []string, fileName string) string {
	for _, value := range changeds {
		if value == fileName {
			return "true"
		}
	}
	return "false"
}

func comparePaths(changeds []string, path string) string {
	for _, value := range changeds {
		if strings.Contains(value, path) {
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
	result := append(changes.HeadCommit.Modified, changes.HeadCommit.Added...)
	compareValue := flag.String("compare", "", "define operation mode")
	flag.Parse()
	if *compareValue != "" {
		fmt.Println(compare(result, *compareValue))
	} else {
		fmt.Println(returnModifieds(result))
	}
}
