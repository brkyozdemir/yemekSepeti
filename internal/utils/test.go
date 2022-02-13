package utils

import (
	"github.com/yemekSepeti/internal/database"
	"github.com/yemekSepeti/internal/repository/concrete"
)

func GetTestDB() database.API {
	db := concrete.NewMemDB()
	api := database.API{
		Db: db,
	}

	return api
}
