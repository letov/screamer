package prof

import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	srv *http.Server
}

func NewProfServer(
	lc fx.Lifecycle,
	log *zap.SugaredLogger,
) *Server {
	router := chi.NewRouter()
	router.Handle("/debug/pprof/*", http.DefaultServeMux)
	srv := &http.Server{Addr: "localhost:8084", Handler: router}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting pprof server: ", srv.Addr)
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
	return &Server{srv}
}
