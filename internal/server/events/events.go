package events

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/server/config"
	"screamer/internal/server/services"
)

type ServerEvents struct {
	events []*event_loop.Event
}

func (a *ServerEvents) PushEvent(e *event_loop.Event) {
	a.events = append(a.events, e)
}

func (a *ServerEvents) GetEvents() []*event_loop.Event {
	return a.events
}

func NewEvents(lc fx.Lifecycle, log *zap.SugaredLogger, c *config.Config, bs *services.BackupService) event_loop.Events {
	es := &ServerEvents{
		events: make([]*event_loop.Event, 0),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if c.Restore && c.StoreInterval > 0 {
				es.PushEvent(event_loop.NewEvent(ctx, "Store backup", c.StoreInterval, bs.Save, log))
			}
			return nil
		},
	})

	return es
}
