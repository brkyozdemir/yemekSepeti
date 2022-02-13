package queries

import (
	"encoding/json"
	"github.com/yemekSepeti/internal/database"
	"net/http"
)

func GetObjectList(api database.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store := api.Db.GetObjectList()

		json.NewEncoder(w).Encode(store)
	}
}
