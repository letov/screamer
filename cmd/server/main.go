package main

import (
	"go.uber.org/fx"
	"net/http"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/server/di"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(func(*event_loop.EventLoop, *http.Server) {}),
	).Run()
}
