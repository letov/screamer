package main

import (
	"screamer/internal/agent/config"
	"screamer/internal/collector"
	"screamer/internal/common/event-loop"
	"screamer/internal/pusher"
)

func init() {
	config.Init()
	collector.Init()
}

func main() {
	c := config.GetConfig()

	event_loop.Run([]*event_loop.Event{
		event_loop.NewEvent(*c.PollInterval, collector.UpdateMetrics),
		event_loop.NewEvent(*c.ReportInterval, pusher.SendData),
	})

	for {
	}
}
