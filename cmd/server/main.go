package main

import (
	"screamer/internal/common/helpers/build"
	"screamer/internal/server/application/app"
	"screamer/internal/server/infrastructure/di"

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
