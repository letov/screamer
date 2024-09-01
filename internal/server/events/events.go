package events

import (
	event_loop "screamer/internal/common/event-loop"
	"screamer/internal/server/config"
	"screamer/internal/server/services"
)

func NewEvents(c *config.Config, bs *services.BackupService) []*event_loop.Event {
	es := make([]*event_loop.Event, 0)

	if c.Restore && c.StoreInterval > 0 {
		es = append(es, event_loop.NewEvent(c.StoreInterval, bs.Save))
	}

	return es
}
