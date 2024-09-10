package main

import (
	"go.uber.org/fx"
	"screamer/internal/agent/di"
	event_loop "screamer/internal/common/eventloop"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(func(*event_loop.EventLoop) {}),
	).Run()
}
