package config

import (
	"github.com/yemekSepeti/config/helpers"
)

type App struct {
	Name string
	Env  string
}

func NewApp() App {
	return App{
		Name: helpers.Getenv("APP_NAME", "yemek-sepeti-in-memory-key-value-storage"),
		Env:  helpers.Getenv("APP_ENV", "local"),
	}
}
