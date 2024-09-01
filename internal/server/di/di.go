package di

import (
	"go.uber.org/dig"
	"screamer/internal/server/config"
	"screamer/internal/server/handlers"
	"screamer/internal/server/repositories"
	"screamer/internal/server/router"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.NewConfig)
	_ = container.Provide(repositories.NewMemoryRepository)

	_ = container.Provide(handlers.NewHomeHandler)
	_ = container.Provide(handlers.NewUpdateMetricHandler)
	_ = container.Provide(handlers.NewValueMetricHandler)
	_ = container.Provide(router.NewRouter)

	return container
}
