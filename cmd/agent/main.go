package main

import (
	"screamer/internal/agent/di"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/common/prof"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(func(*event_loop.EventLoop, *prof.Server) {}),
	).Run()
}
