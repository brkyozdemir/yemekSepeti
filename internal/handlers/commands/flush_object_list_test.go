package commands

import (
	"github.com/yemekSepeti/internal/utils"
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
