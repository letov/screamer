package main

import (
	"net/http"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/server/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(func(*event_loop.EventLoop, *http.Server) {}),
	).Run()
}
