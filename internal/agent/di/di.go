package di

import (
	"go.uber.org/dig"
	"screamer/internal/agent/config"
	"screamer/internal/agent/events"
	metric_sources "screamer/internal/agent/metric-sources"
	"screamer/internal/agent/repositories"
	"screamer/internal/agent/services"
	event_loop "screamer/internal/common/event-loop"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.NewConfig)
	_ = container.Provide(repositories.NewMemoryRepository)
	_ = container.Provide(metric_sources.NewMetricSources)

	_ = container.Provide(services.NewShutdownService)
	_ = container.Provide(services.NewProcessingService)
	_ = container.Provide(services.NewSendingService)

	_ = container.Provide(event_loop.NewEventLoop)
	_ = container.Provide(events.NewEvents)

	return container
}
