package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"screamer/internal/server/config"
	"screamer/internal/server/handlers"
	"screamer/internal/server/middlewares"
	"time"
)

type Router struct {
	config        *config.Config
	homeHandler   *handlers.HomeHandler
	updateHandler *handlers.UpdateMetricHandler
	valueHandler  *handlers.ValueMetricHandler
}

func (rtr *Router) Run() {
	router := rtr.GetRouter()

	addr := fmt.Sprintf(":%v", rtr.config.NetAddress.Port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

func (rtr *Router) RunAsync() {
	go rtr.Run()
}

func (rtr *Router) GetRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.Logger)
	r.Use(middlewares.Compress([]string{"application/json", "text/html"}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", rtr.homeHandler.GetHandler())

	r.Route("/update", func(r chi.Router) {
		r.Post("/", rtr.updateHandler.GetHandlerJson())
		r.Post("/{type:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}/{value}", rtr.updateHandler.GetHandlerParams())
	})

	r.Route("/value", func(r chi.Router) {
		r.Post("/", rtr.valueHandler.GetHandlerJson())
		r.Get("/{type:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}", rtr.valueHandler.GetHandlerParams())
	})

	return r
}

func NewRouter(c *config.Config, hh *handlers.HomeHandler, uh *handlers.UpdateMetricHandler, vh *handlers.ValueMetricHandler) *Router {
	return &Router{
		config:        c,
		homeHandler:   hh,
		updateHandler: uh,
		valueHandler:  vh,
	}
}
