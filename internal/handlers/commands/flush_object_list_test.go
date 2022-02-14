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

func TestFlushObjectList(t *testing.T) {
	api := utils.GetTestDB()

	key := "key"
	value := "value"

	api.Db.CreateObject(key, value)

	if len(api.Db.GetObjectList()) > 0 {
		t.Errorf("Couldn't flush the memory!")
	}
}

func TestIntegrationFlushObjectList(t *testing.T) {
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

	createReq, _ := http.NewRequest("POST", url+"/api/keys", responseBody)
	_, _ = client.Do(createReq)

	res, err := http.NewRequest(http.MethodDelete, url+"/api/keys", nil)
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

	if resMap["message"] != "Store flushed successfully!" {
		t.Errorf("It is not flushed!")
	}
}
