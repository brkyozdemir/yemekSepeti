package utils

import (
	"github.com/yemekSepeti/internal/database"
	"github.com/yemekSepeti/internal/repository/concrete"
	"net/http"
)

func GetTestDB() database.API {
	db := concrete.NewMemDB()
	api := database.API{
		Db: db,
	}

	return api
}

func Refresh(client *http.Client) {
	refresh, _ := http.NewRequest(http.MethodDelete, "http://localhost:9000/api/keys", nil)
	_, _ = client.Do(refresh)
}
