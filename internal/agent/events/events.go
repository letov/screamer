package events

import (
	"screamer/internal/agent/config"
	"screamer/internal/agent/services"
	event_loop "screamer/internal/common/eventloop"
)

func NewEvents(c *config.Config, ps *services.ProcessingService, ss *services.SendingService) []*event_loop.Event {
	return []*event_loop.Event{
		event_loop.NewEvent(c.PollInterval, ps.UpdateMetrics),
		event_loop.NewEvent(c.ReportInterval, ss.SendMetrics),
	}
}
