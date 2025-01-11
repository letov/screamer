package mux

import (
	"screamer/internal/server/infrastructure/config"
	handlers2 "screamer/internal/server/infrastructure/http/handlers"
	"screamer/internal/server/infrastructure/http/middlewares"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewMux(
	c *config.Config,
	hh *handlers2.HomeHandler,
	uh *handlers2.UpdateMetricHandler,
	ush *handlers2.UpdateBatchMetricHandler,
	uoh *handlers2.UpdateMetricOldHandler,
	vh *handlers2.ValueMetricHandler,
	voh *handlers2.ValueMetricOldHandler,
	ph *handlers2.PingHandler,
	log *zap.SugaredLogger,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Compress(5, "application/json", "text/html"))
	r.Use(middleware.Timeout(10 * time.Second))

	r.Use(middlewares.Logger)
	r.Use(middlewares.CheckHash(c))
	r.Use(middlewares.Decrypt(c, log))
	r.Use(middlewares.TrustedSubnet(c))
	r.Use(middlewares.Curl)

	r.Get("/", hh.Handler)
	r.Get("/ping", ph.Handler)

	r.Route("/update", func(r chi.Router) {
		r.Post("/", uh.Handler)
		r.Post("/{type:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}/{value}", uoh.Handler)
	})

	r.Post("/updates", ush.Handler)

	r.Route("/value", func(r chi.Router) {
		r.Post("/", vh.Handler)
		r.Get("/{type:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}", voh.Handler)
	})

	return r
}
