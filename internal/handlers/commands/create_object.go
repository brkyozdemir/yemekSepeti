package commands

import (
	"encoding/json"
	"github.com/yemekSepeti/internal/database"
	"log"
	"net/http"
)

func CreateObject(api database.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body map[string]string
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			log.Fatal(err)
		}

		var object map[string]string
		for key, value := range body {
			object = api.Db.CreateObject(key, value)
		}

		if object != nil {
			w.WriteHeader(http.StatusCreated)
		}

		json.NewEncoder(w).Encode(object)
	}
}
