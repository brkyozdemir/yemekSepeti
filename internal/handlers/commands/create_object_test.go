package commands

import (
	"bytes"
	"encoding/json"
	"github.com/yemekSepeti/config"
	"github.com/yemekSepeti/internal/utils"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateObject(t *testing.T) {
	api := utils.GetTestDB()

	key := "key"
	value := "value"

	object := api.Db.CreateObject(key, value)

	for k, v := range object {
		if k != key {
			t.Errorf("Key is not true.")
		}

		if v != value {
			t.Errorf("Value is not true.")
		}
	}
}

func TestIntegrationCreateObject(t *testing.T) {
	client := &http.Client{}
	utils.Refresh(client)

	key := "key"
	value := "value"

	postBody, _ := json.Marshal(map[string]string{
		key: value,
	})
	responseBody := bytes.NewBuffer(postBody)

	cfg := config.GetConfig()
	url := cfg.App.BaseUrl + ":" + cfg.App.Port

	res, err := http.NewRequest("POST", url+"/api/keys", responseBody)
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
