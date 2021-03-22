package scheduler

import (
	"time"

	"github.com/shikhar1996/Covid19/src/database"
	"go.uber.org/zap"
)

func ScheduleUpdateDatabase() {
	// Updating the database every hour
	for t := range time.NewTicker(3600 * time.Second).C {
		zap.String("Scheduler: ", t.String())
		// fmt.Println(t)
		data, err := database.Getdata()
		if err != nil {
			zap.String("Error: Database Updation", err.Error())
		} else {
			err = database.Updatedata(data)
			if err != nil {
				zap.String("Error: Database Updation", err.Error())
			}
		}
	}
}
