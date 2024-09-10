package eventloop

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

type Event struct {
	Name     string
	Duration time.Duration
	Event    func()
}

type EventLoop struct {
	events []*Event
	log    *zap.SugaredLogger
}

func intToSecond(i int) time.Duration {
	return time.Duration(int64(i)) * time.Second
}

func NewEvent(n string, s int, e func()) *Event {
	return &Event{
		Name:     n,
		Duration: intToSecond(s),
		Event:    e,
	}
}

func (el *EventLoop) Run() {
	for _, event := range el.events {
		ticker := time.NewTicker(event.Duration)
		event := event
		go func() {
			for {
				if _, ok := <-ticker.C; ok {
					el.log.Info("Run event: ", event.Name)
					event.Event()
				}
			}
		}()
	}
}

func NewEventLoop(lc fx.Lifecycle, log *zap.SugaredLogger, es []*Event) *EventLoop {
	el := &EventLoop{
		events: es,
		log:    log,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting event loop")
			el.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping event loop")
			return nil
		},
	})

	return el
}
