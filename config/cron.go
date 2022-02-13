package config

import "github.com/yemekSepeti/config/helpers"

type Cron struct {
	TimeIntervalInMinutes string
}

func NewCron() Cron {
	return Cron{
		TimeIntervalInMinutes: helpers.Getenv("TIME_INTERVAL_IN_MINUTES", "1"),
	}
}
