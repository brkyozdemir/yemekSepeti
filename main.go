package main

import (
	"github.com/yemekSepeti/internal/crons"
	"github.com/yemekSepeti/internal/database"
	"github.com/yemekSepeti/internal/repository/concrete"
	"github.com/yemekSepeti/internal/routes"
	"github.com/yemekSepeti/internal/utils"
)

func main() {
	db := concrete.NewMemDB()

	store := utils.ReadFromResources()
	db.Store = store

	api := database.API{
		Db: db,
	}

	go crons.WriteFile(api)

	rts := routes.Routes(api)
	routes.Start(rts)
}
