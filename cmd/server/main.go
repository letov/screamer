package main

import (
	"screamer/internal/grab"
	"screamer/internal/logger"
	"screamer/internal/server/config"
	"screamer/internal/storage"
)

func init() {
	config.Init()
	storage.Init()
	logger.Init()
	go grab.Init()
}

func main() {
	defer stopped()
	for {

	}

	//c := config.GetConfigA()
	//
	//event_loop.Run([]*event_loop.Event{
	//	event_loop.NewEvent(c.PollInterval, collector.UpdateMetrics),
	//	event_loop.NewEvent(c.ReportInterval, pusher.SendData),
	//})
}

func stopped() {

}
