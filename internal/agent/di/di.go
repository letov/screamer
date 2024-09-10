package di

import (
	"go.uber.org/fx"
	"screamer/internal/agent/config"
	"screamer/internal/agent/events"
	metric_sources "screamer/internal/agent/metricsources"
	"screamer/internal/agent/repositories"
	"screamer/internal/agent/services"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/common/logger"
	"screamer/internal/server/handlers"
)

func InjectApp() fx.Option {
	return fx.Provide(
		config.NewConfig,
		logger.NewLogger,

		repositories.NewMemoryRepository,
		metric_sources.NewMetricSources,

		services.NewProcessingService,
		services.NewSendingService,

		handlers.NewHomeHandler,
		handlers.NewUpdateMetricHandler,
		handlers.NewUpdateMetricOldHandler,
		handlers.NewValueMetricHandler,
		handlers.NewValueMetricOldHandler,

		event_loop.NewEventLoop,
		events.NewEvents,
	)
}
