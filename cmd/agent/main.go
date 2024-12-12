package main

import (
	"screamer/internal/agent/di"
	"screamer/internal/common/build"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/common/prof"

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
		fx.Invoke(func(*event_loop.EventLoop, *prof.Server) {}),
	).Run()
}
