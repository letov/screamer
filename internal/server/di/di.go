package di

import (
	"go.uber.org/dig"
	event_loop "screamer/internal/common/event-loop"
	"screamer/internal/server/config"
	"screamer/internal/server/events"
	"screamer/internal/server/handlers"
	"screamer/internal/server/repositories"
	"screamer/internal/server/router"
	"screamer/internal/server/services"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.NewConfig)
	_ = container.Provide(repositories.NewMemoryRepository)

	_ = container.Provide(services.NewBackupService)
	_ = container.Provide(services.NewLoadService)
	_ = container.Provide(services.NewShutdownService)

	_ = container.Provide(handlers.NewHomeHandler)
	_ = container.Provide(handlers.NewUpdateMetricHandler)
	_ = container.Provide(handlers.NewValueMetricHandler)
	_ = container.Provide(router.NewRouter)

	_ = container.Provide(event_loop.NewEventLoop)
	_ = container.Provide(events.NewEvents)

	return container
}
