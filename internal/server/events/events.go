package events

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
)

type ServerEvents struct {
	events []*event_loop.Event
}

func (se *ServerEvents) PushEvent(e *event_loop.Event) {
	se.events = append(se.events, e)
}

func (se *ServerEvents) GetEvents() []*event_loop.Event {
	return se.events
}

func NewEvents(lc fx.Lifecycle, log *zap.SugaredLogger, c *config.Config, fr *repositories.FileRepository) event_loop.Events {
	es := &ServerEvents{
		events: make([]*event_loop.Event, 0),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if c.Restore && c.StoreInterval > 0 {
				es.PushEvent(event_loop.NewEvent("Store backup", c.StoreInterval, fr.SaveAllToFile, log))
			}
			return nil
		},
	})

	return es
}
