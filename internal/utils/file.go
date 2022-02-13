package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func WriteFile(data interface{}) {
	file, _ := json.MarshalIndent(data, "", " ")
	currentTime := time.Now()
	today := currentTime.Format("20060102")
	_ = ioutil.WriteFile("resources/"+today+"-data.json", file, 0644)
}

func ReadFromResources() []map[string]string {
	currentTime := time.Now()
	today := currentTime.Format("20060102")

	fileName := today + "-data.json"
	jsonFile, err := os.Open("resources/" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var store []map[string]string

	json.Unmarshal(byteValue, &store)

	return store
}
