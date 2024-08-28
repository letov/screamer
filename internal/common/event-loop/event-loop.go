package event_loop

import (
	"time"
)

type Event struct {
	Duration time.Duration
	Event    func()
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

func Run(events []*Event) {
	for _, event := range events {
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
