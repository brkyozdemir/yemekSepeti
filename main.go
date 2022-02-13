package main

import (
	"github.com/yemekSepeti/internal/database"
	"github.com/yemekSepeti/internal/repository/concrete"
	"github.com/yemekSepeti/internal/routes"
)

func main() {
	db := concrete.NewMemDB()

	api := database.API{
		Db: db,
	}

	rts := routes.Routes(api)
	routes.Start(rts)
}
