package commands

import (
	"encoding/json"
	"github.com/yemekSepeti/internal/database"
	"net/http"
)

func FlushObjectList(api database.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, _ := api.Db.FlushObjectList()

		var message map[string]string
		if result {
			message = map[string]string{"message": "Store flushed successfully!"}
		} else {
			message = map[string]string{"message": "An Error occured!"}
		}

		json.NewEncoder(w).Encode(message)
	}
}
