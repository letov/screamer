package di

import (
	event_loop "screamer/internal/common/infrastructure/eventloop"
	"screamer/internal/common/infrastructure/grpcclient"
	"screamer/internal/common/infrastructure/logger"
	"screamer/internal/server/application/events"
	"screamer/internal/server/application/services"
	config2 "screamer/internal/server/infrastructure/config"
	"screamer/internal/server/infrastructure/db"
	"screamer/internal/server/infrastructure/grpc/grpcserver"
	handlers2 "screamer/internal/server/infrastructure/http/handlers"
	"screamer/internal/server/infrastructure/http/httpserver"
	"screamer/internal/server/infrastructure/http/mux"
	repositories2 "screamer/internal/server/infrastructure/repositories"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func InjectApp() fx.Option {
	return fx.Provide(
		config2.NewConfig,
		logger.NewLogger,
		db.NewDB,

		repositories2.NewDBRepository,
		repositories2.NewFileRepository,
		repositories2.NewMemoryRepository,

		func(
			c *config2.Config,
			log *zap.SugaredLogger,
			db *repositories2.DBRepository,
			fr *repositories2.FileRepository,
			mr *repositories2.MemoryRepository,
		) repositories2.Repository {
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

		handlers2.NewHomeHandler,
		handlers2.NewUpdateMetricHandler,
		handlers2.NewUpdateBatchMetricHandler,
		handlers2.NewUpdateMetricOldHandler,
		handlers2.NewValueMetricHandler,
		handlers2.NewValueMetricOldHandler,
		handlers2.NewPingHandler,

		event_loop.NewEventLoop,
		events.NewEvents,

		mux.NewMux,
		httpserver.NewHTTPServer,
		grpcserver.NewGRPCServer,
		grpcclient.NewGRPCClient,
	)
}
