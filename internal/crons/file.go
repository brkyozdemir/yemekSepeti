package crons

import (
	"fmt"
	"github.com/yemekSepeti/config"
	"github.com/yemekSepeti/internal/database"
	"github.com/yemekSepeti/internal/utils"
	"strconv"
	"time"
)

type jobTicker struct {
	timer *time.Timer
}

func WriteFile(api database.API) {
	jobTicker := &jobTicker{}
	jobTicker.updateTimer()
	for {
		<-jobTicker.timer.C
		fmt.Println("File will be created")
		currentTime := time.Now()
		today := currentTime.Format("20060102")
		fileName := "resources/" + today + "-data.json"
		utils.WriteFile(api.Db.GetObjectList(), fileName)
		fmt.Println("File created")
		jobTicker.updateTimer()
	}
}

func (t *jobTicker) updateTimer() {
	cfg := config.GetConfig()
	minutes, _ := strconv.Atoi(cfg.Cron.TimeIntervalInMinutes)

	diff := time.Minute * time.Duration(minutes)
	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}
