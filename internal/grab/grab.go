package grab

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"screamer/internal/config"
	"screamer/internal/grab/handlers"
	"screamer/internal/grab/middlewares"
	"time"
)

func Init() {
	c := config.GetConfigS()
	router := getRouter()

	addr := fmt.Sprintf(":%v", c.Port)
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

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handlers.Home)

	r.Post("/update", handlers.UpdateMetric)

	r.Post("/value", handlers.ValueMetric)

	return r
}
