package mux

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"screamer/internal/server/handlers"
	"screamer/internal/server/middlewares"
	"time"
)

func NewMux(
	hh *handlers.HomeHandler,
	uh *handlers.UpdateMetricHandler,
	uoh *handlers.UpdateMetricOldHandler,
	vh *handlers.ValueMetricHandler,
	voh *handlers.ValueMetricOldHandler,
	ph *handlers.PingHandler,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "application/json", "text/html"))
	r.Use(middlewares.Logger)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", hh.Handler)
	r.Get("/ping", ph.Handler)

	r.Route("/update", func(r chi.Router) {
		r.Post("/", uh.Handler)
		r.Post("/{type:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}/{value}", uoh.Handler)
	})

	r.Route("/value", func(r chi.Router) {
		r.Post("/", vh.Handler)
		r.Get("/{type:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}", voh.Handler)
	})

	return r
}
