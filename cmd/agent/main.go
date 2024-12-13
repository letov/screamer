package main

import (
	"screamer/internal/agent/app"
	"screamer/internal/agent/di"
	"screamer/internal/common/build"

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
		fx.Invoke(app.Start),
	).Run()
}
