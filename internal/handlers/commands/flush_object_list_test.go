package commands

import (
	"encoding/json"
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
	api := utils.GetTestDB()

	key := "key"
	value := "value"

	api.Db.CreateObject(key, value)

	res, err := http.NewRequest(http.MethodDelete, "http://localhost:9000/api/keys", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	client := &http.Client{}
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
