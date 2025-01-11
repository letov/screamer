package di

import (
	"screamer/internal/agent/application/agentservices"
	"screamer/internal/agent/infrastructure/config"
	"screamer/internal/agent/infrastructure/events"
	metric_sources "screamer/internal/agent/infrastructure/metricsources"
	"screamer/internal/agent/infrastructure/store"
	event_loop "screamer/internal/common/infrastructure/eventloop"
	"screamer/internal/common/infrastructure/grpcclient"
	"screamer/internal/common/infrastructure/logger"
	"screamer/internal/common/infrastructure/prof"
	"screamer/internal/server/infrastructure/http/handlers"

	"go.uber.org/fx"
)

func InjectApp() fx.Option {
	return fx.Provide(
		config.NewConfig,
		logger.NewLogger,

		store.NewMemoryRepository,
		metric_sources.NewMetricSources,

		agentservices.NewProcessing,
		agentservices.NewSending,

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
