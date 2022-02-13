package config

import "sync"

var once sync.Once
var config *Config

type Config struct {
	App  App
	Cron Cron
}

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			App:  NewApp(),
			Cron: NewCron(),
		}
	})

	return config
}
