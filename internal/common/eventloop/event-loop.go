package eventloop

import (
	"time"
)

type Event struct {
	Duration time.Duration
	Event    func()
}

type EventLoop struct {
	Events []*Event
}

func intToSecond(i int) time.Duration {
	return time.Duration(int64(i)) * time.Second
}

func NewEvent(s int, e func()) *Event {
	return &Event{
		Duration: intToSecond(s),
		Event:    e,
	}
}

func NewEventLoop(es []*Event) *EventLoop {
	return &EventLoop{
		Events: es,
	}
}

func (el *EventLoop) Run() {
	for _, event := range el.Events {
		ticker := time.NewTicker(event.Duration)
		event := event
		go func() {
			for {
				if _, ok := <-ticker.C; ok {
					event.Event()
				}
			}
		}()
	}
}
