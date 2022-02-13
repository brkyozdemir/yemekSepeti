package queries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yemekSepeti/internal/utils"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetObjectByKey(t *testing.T) {
	api := utils.GetTestDB()

	key := "key"
	value := "value"

	api.Db.CreateObject(key, value)

	object := api.Db.GetObjectByKey(key)

	for k, v := range object {
		if k != key {
			t.Errorf("Key is not true.")
		}

		if v != value {
			t.Errorf("Value is not true.")
		}
	}
}

func TestGetObjectByNotExistingKey(t *testing.T) {
	api := utils.GetTestDB()

	object := api.Db.GetObjectByKey("key")

	for k, v := range object {
		// Bad usage of magic strings
		if k != "message" {
			t.Errorf("Key is not true.")
		}

		if v != "Key not found!" {
			t.Errorf("Value is not true.")
		}
	}
}

func TestIntegrationGetObjectByKey(t *testing.T) {
	client := &http.Client{}
	utils.Refresh(client)

	key := "key"
	value := "value"
	postBody, _ := json.Marshal(map[string]string{
		key: value,
	})
	responseBody := bytes.NewBuffer(postBody)

	createReq, _ := http.NewRequest("POST", "http://localhost:9000/api/keys", responseBody)
	_, _ = client.Do(createReq)

	res, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:9000/api/keys/%s", key), nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp, err := client.Do(res)
	if err != nil {
		t.Errorf(err.Error())
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	var resMap map[string]string
	json.Unmarshal(resBody, &resMap)

	for k, v := range resMap {
		if k != key {
			t.Errorf("Key is not true")
		}

		if v != value {
			t.Errorf("Value is not true")
		}
	}
}
