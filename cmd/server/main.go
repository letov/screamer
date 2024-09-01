package main

import (
	event_loop "screamer/internal/common/event-loop"
	infinity_loop "screamer/internal/common/infinity-loop"
	"screamer/internal/server/di"
	"screamer/internal/server/router"
	"screamer/internal/server/services"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(router *router.Router, el *event_loop.EventLoop, ls *services.LoadService, ss *services.ShutdownService) {
		ls.Run()

		router.RunAsync()
		el.Run()

		ss.Run()
	})

	if err != nil {
		panic(err)
	}

	infinity_loop.Run()
}
