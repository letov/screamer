package events

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"screamer/internal/agent/config"
	"screamer/internal/agent/services"
	event_loop "screamer/internal/common/eventloop"
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

func NewEvents(lc fx.Lifecycle, log *zap.SugaredLogger, c *config.Config, ps *services.ProcessingService, ss *services.SendingService) event_loop.Events {
	es := &AgentEvents{
		events: make([]*event_loop.Event, 0),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			es.PushEvent(event_loop.NewEvent(ctx, "UpdateMetrics", c.PollInterval, ps.UpdateMetrics, log))
			es.PushEvent(event_loop.NewEvent(ctx, "SendMetrics", c.ReportInterval, ss.SendMetrics, log))
			return nil
		},
	})

	return es
}
