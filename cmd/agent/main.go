package main

import (
	"screamer/internal/agent/di"
	event_loop "screamer/internal/common/event-loop"
	infinity_loop "screamer/internal/common/infinity-loop"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(el *event_loop.EventLoop) {
		el.Run()
	})

	if err != nil {
		panic(err)
	}

	infinity_loop.Run()
}
