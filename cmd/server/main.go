package main

import (
	"screamer/internal/collector"
	event_loop "screamer/internal/common/event-loop"
	"screamer/internal/config"
	"screamer/internal/grab"
	"screamer/internal/logger"
	"screamer/internal/pusher"
	"screamer/internal/storage"
)

func init() {
	config.InitServer()
	storage.Init()
	logger.Init()
	go grab.Init()
}

func main() {
	defer stopped()

	c := config.GetConfigA()

	event_loop.Run([]*event_loop.Event{
		event_loop.NewEvent(c.PollInterval, collector.UpdateMetrics),
		event_loop.NewEvent(c.ReportInterval, pusher.SendData),
	})
}

func stopped() {

}
