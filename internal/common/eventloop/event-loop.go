package eventloop

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

type FuncWithCtx = func(ctx context.Context)

type Event struct {
	Name     string
	Duration time.Duration
	Event    FuncWithCtx
	Cancel   func()
}

type Events interface {
	PushEvent(e *Event)
	GetEvents() []*Event
}

type EventLoop struct {
	events Events
	log    *zap.SugaredLogger
}

func intToSecond(i int) time.Duration {
	return time.Duration(int64(i)) * time.Second
}

func (e *Event) SetCancel(c func()) {
	e.Cancel = c
}

func (e *Event) CallCancel() {
	if e.Cancel != nil {
		e.Cancel()
	}
}

func NewEvent(n string, s int, e FuncWithCtx, log *zap.SugaredLogger) *Event {
	log.Info("Registered new event: ", n, " on every ", s, " sec")

	return &Event{
		Name:     n,
		Duration: intToSecond(s),
		Event:    e,
		Cancel:   nil,
	}
}

func (el *EventLoop) Run() {
	for _, event := range el.events.GetEvents() {
		ticker := time.NewTicker(event.Duration)
		event := event
		go func() {
			for {
				if _, ok := <-ticker.C; ok {
					el.log.Info("Run event: ", event.Name)
					event.CallCancel()
					ctxTimeout, cancel := context.WithTimeout(context.Background(), event.Duration)
					event.SetCancel(cancel)
					event.Event(ctxTimeout)
				}
			}
		}()
	}
}

func NewEventLoop(lc fx.Lifecycle, log *zap.SugaredLogger, es Events) *EventLoop {
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
			for _, event := range el.events.GetEvents() {
				event.CallCancel()
			}
			return nil
		},
	})

	return el
}
