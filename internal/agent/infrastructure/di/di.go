package di

import (
	"screamer/internal/agent/application/events"
	services2 "screamer/internal/agent/application/services"
	"screamer/internal/agent/infrastructure/config"
	metric_sources "screamer/internal/agent/infrastructure/metricsources"
	"screamer/internal/agent/infrastructure/repositories"
	event_loop "screamer/internal/common/infrastructure/eventloop"
	"screamer/internal/common/infrastructure/grpcclient"
	"screamer/internal/common/infrastructure/logger"
	"screamer/internal/common/infrastructure/prof"
	"screamer/internal/server/handlers"

	"go.uber.org/fx"
)

func InjectApp() fx.Option {
	return fx.Provide(
		config.NewConfig,
		logger.NewLogger,

		repositories.NewMemoryRepository,
		metric_sources.NewMetricSources,

		services2.NewProcessing,
		services2.NewSending,

		handlers.NewHomeHandler,
		handlers.NewUpdateMetricHandler,
		handlers.NewUpdateMetricOldHandler,
		handlers.NewValueMetricHandler,
		handlers.NewValueMetricOldHandler,

		events.NewEvents,
		event_loop.NewEventLoop,

		prof.NewProfServer,
		grpcclient.NewGRPCClient,
	)
}
