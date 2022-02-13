package queries

import (
	"encoding/json"
	"github.com/yemekSepeti/internal/database"
	"github.com/yemekSepeti/internal/utils"
	"net/http"
)

func GetObjectByKey(api database.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := utils.GetField(r, 0)

		obj := api.Db.GetObjectByKey(key)

		json.NewEncoder(w).Encode(obj)
	}
}
