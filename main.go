package main

import (
	"github.com/shikhar1996/Covid19/src/scheduler"
	"github.com/shikhar1996/Covid19/src/server"
)

func main() {

	// Start logging
	// logger.Init()
	// Start the server
	go server.Redirect()

	// Start scheduler
	scheduler.ScheduleUpdateDatabase()

}
