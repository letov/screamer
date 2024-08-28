package grab

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"screamer/internal/grab/handlers"
	"screamer/internal/grab/middlewares"
	"screamer/internal/server/config"
	"time"
)

func Init() {
	c := config.GetConfig()
	router := getRouter()

	addr := fmt.Sprintf(":%v", (*c).NetAddress.Port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

func getRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.Logger)
	r.Use(middlewares.Compress([]string{"application/json", "text/html"}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handlers.Home)

	r.Route("/update", func(r chi.Router) {
		r.Post("/", handlers.UpdateMetric)
		r.Post("/{label:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}/{value}", handlers.UpdateMetricOld)
	})

	r.Route("/value", func(r chi.Router) {
		r.Post("/", handlers.ValueMetric)
		r.Get("/{label:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}", handlers.ValueMetricOld)
	})

	return r
}
