package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadYTJson(filename string) []YTJson {
	var YTJsons []YTJson
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	error := json.Unmarshal(byteValue, &YTJsons)
	if error != nil {
		fmt.Println(error)
	}
	return YTJsons
}
