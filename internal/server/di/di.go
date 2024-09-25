package di

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
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
		db.NewDB,

		repositories.NewDBRepository,
		repositories.NewFileRepository,
		repositories.NewMemoryRepository,

		func(
			c *config.Config,
			log *zap.SugaredLogger,
			db *repositories.DBRepository,
			fr *repositories.FileRepository,
			mr *repositories.MemoryRepository,
		) repositories.Repository {
			switch true {
			case len(c.DBAddress) > 0:
				log.Info("DB as repo source")
				return db
			case c.Restore:
				log.Info("FILE as repo source")
				return fr
			default:
				log.Info("Memory as repo source")
				return mr
			}
		},

		services.NewMetricService,

		handlers.NewHomeHandler,
		handlers.NewUpdateMetricHandler,
		handlers.NewUpdatesMetricHandler,
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
