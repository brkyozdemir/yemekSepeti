package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func WriteFile(data interface{}, fileName string) {
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(fileName, file, 0644)
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
