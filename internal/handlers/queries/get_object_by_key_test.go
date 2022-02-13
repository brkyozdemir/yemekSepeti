package queries

import (
	"github.com/yemekSepeti/internal/utils"
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
