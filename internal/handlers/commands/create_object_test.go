package commands

import (
	"github.com/yemekSepeti/internal/utils"
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
