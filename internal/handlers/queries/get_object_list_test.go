package queries

import (
	"github.com/yemekSepeti/internal/utils"
	"testing"
)

func TestGetObjectList(t *testing.T) {
	api := utils.GetTestDB()

	value := "value"

	for i := 0; i < 10; i++ {
		api.Db.CreateObject(string(rune(i)), value)
	}

	if len(api.Db.GetObjectList()) != 10 {
		t.Errorf("List must be 10 items long!")
	}
}
