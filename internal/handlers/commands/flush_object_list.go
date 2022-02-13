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
			// Bad usage of magic strings
			message = map[string]string{"message": "Store flushed successfully!"}
		} else {
			// Bad usage of magic strings
			message = map[string]string{"message": "An Error occured!"}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	}
}
