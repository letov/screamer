package di

import (
	"go.uber.org/fx"
	event_loop "screamer/internal/common/eventloop"
	"screamer/internal/common/logger"
	"screamer/internal/server/config"
	"screamer/internal/server/db"
	"screamer/internal/server/events"
	"screamer/internal/server/handlers"
	"screamer/internal/server/httpserver"
	"screamer/internal/server/mux"
	"screamer/internal/server/repositories"
	"screamer/internal/server/services"
)

func InjectApp() fx.Option {
	return fx.Provide(
		config.NewConfig,
		logger.NewLogger,

		repositories.NewMemoryRepository,
		db.NewDB,

		services.NewMetricService,
		services.NewBackupService,

		handlers.NewHomeHandler,
		handlers.NewUpdateMetricHandler,
		handlers.NewUpdateMetricOldHandler,
		handlers.NewValueMetricHandler,
		handlers.NewValueMetricOldHandler,
		handlers.NewPingHandler,

		event_loop.NewEventLoop,
		events.NewEvents,

		mux.NewMux,
		httpserver.NewHTTPServer,
	)
}
