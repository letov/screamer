package httpserver

import (
	"context"
	"net"
	"net/http"
	"screamer/internal/server/config"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, mux *chi.Mux, log *zap.SugaredLogger, c *config.Config) *http.Server {
	srv := &http.Server{Addr: c.NetAddress.String(), Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server: ", srv.Addr)
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					log.Error(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
