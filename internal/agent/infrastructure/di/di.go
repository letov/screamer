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
	handlers2 "screamer/internal/server/infrastructure/http/handlers"

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

		handlers2.NewHomeHandler,
		handlers2.NewUpdateMetricHandler,
		handlers2.NewUpdateMetricOldHandler,
		handlers2.NewValueMetricHandler,
		handlers2.NewValueMetricOldHandler,

		events.NewEvents,
		event_loop.NewEventLoop,

		prof.NewProfServer,
		grpcclient.NewGRPCClient,
	)
}
