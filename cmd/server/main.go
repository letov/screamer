package main

import (
	infinity_loop "screamer/internal/common/infinity-loop"
	"screamer/internal/server/di"
	"screamer/internal/server/router"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(router *router.Router) {
		router.Run()
	})

	if err != nil {
		panic(err)
	}

	infinity_loop.Run()
}
