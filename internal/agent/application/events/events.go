package events

import (
	"context"
	services2 "screamer/internal/agent/application/services"
	"screamer/internal/agent/infrastructure/config"
	event_loop "screamer/internal/common/infrastructure/eventloop"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type AgentEvents struct {
	events []*event_loop.Event
}

func (a *AgentEvents) PushEvent(e *event_loop.Event) {
	a.events = append(a.events, e)
}

func (a *AgentEvents) GetEvents() []*event_loop.Event {
	return a.events
}

func NewEvents(lc fx.Lifecycle, log *zap.SugaredLogger, c *config.Config, ps *services2.Processing, ss *services2.Sending) event_loop.Events {
	es := &AgentEvents{
		events: make([]*event_loop.Event, 0),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			es.PushEvent(event_loop.NewEvent("UpdateMetrics", c.PollInterval, ps.UpdateMetrics, log))
			es.PushEvent(event_loop.NewEvent("SendMetrics", c.ReportInterval, ss.SendMetrics, log))
			return nil
		},
	})

	return es
}
