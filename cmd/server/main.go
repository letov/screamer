package main

import (
	"go.uber.org/fx"
	"net/http"
	"screamer/internal/server/di"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
