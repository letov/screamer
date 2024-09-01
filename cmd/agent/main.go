package main

import (
	"screamer/internal/agent/di"
	"screamer/internal/agent/services"
	event_loop "screamer/internal/common/event-loop"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(el *event_loop.EventLoop, ss *services.ShutdownService) {
		el.Run()
		ss.Run()
	})

	if err != nil {
		panic(err)
	}
}
