package main

import (
	"screamer/internal/collector"
	"screamer/internal/common/event-loop"
	"screamer/internal/config"
	"screamer/internal/pusher"
)

func init() {
	config.InitAgent()
	collector.Init()
}

func main() {
	c := config.GetConfigA()

	event_loop.Run([]*event_loop.Event{
		event_loop.NewEvent(c.PollInterval, collector.UpdateMetrics),
		event_loop.NewEvent(c.ReportInterval, pusher.SendData),
	})

	for {
	}
}
