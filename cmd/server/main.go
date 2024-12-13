package main

import (
	"screamer/internal/common/build"
	"screamer/internal/server/app"
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
		fx.Invoke(app.Start),
	).Run()
}
