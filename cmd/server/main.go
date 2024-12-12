package main

import (
	"net/http"
	"screamer/internal/common/build"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/server/di"

	"go.uber.org/fx"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func main() {
	build.ShowBuildParams(buildVersion, buildDate, buildCommit)
	fx.New(
		di.InjectApp(),
		fx.Invoke(func(*event_loop.EventLoop, *http.Server) {}),
	).Run()
}
