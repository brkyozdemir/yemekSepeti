package queries

import (
	"bytes"
	"encoding/json"
	"github.com/yemekSepeti/internal/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestGetObjectList(t *testing.T) {
	api := utils.GetTestDB()

	value := "value"

	for i := 0; i < 10; i++ {
		index := strconv.Itoa(i)
		key := "key" + index
		api.Db.CreateObject(key, value)
	}

	if len(api.Db.GetObjectList()) != 10 {
		t.Errorf("List must be 10 items long!")
	}
}

func TestIntegrationGetObjectList(t *testing.T) {
	client := &http.Client{}
	utils.Refresh(client)

	value := "value"
	for i := 0; i < 10; i++ {
		index := strconv.Itoa(i)
		key := "key" + index
		postBody, _ := json.Marshal(map[string]string{
			key: value,
		})
		responseBody := bytes.NewBuffer(postBody)

		createReq, _ := http.NewRequest("POST", "http://localhost:9000/api/keys", responseBody)

		_, _ = client.Do(createReq)
	}

	res, err := http.NewRequest("GET", "http://localhost:9000/api/keys", nil)
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

	var resMap []map[string]string
	json.Unmarshal(resBody, &resMap)
	if len(resMap) != 10 {
		t.Errorf("List must be 10 items long!")
	}
}
