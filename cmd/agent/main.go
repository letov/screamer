package main

import (
	"screamer/internal/agent/application/app"
	"screamer/internal/agent/infrastructure/di"
	"screamer/internal/common/helpers/build"

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
