package grab

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"screamer/internal/config"
	handlers "screamer/internal/grab/handlers"
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
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handlers.Home)

	r.Route("/update", func(r chi.Router) {
		r.Post("/{label:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}/{value}", handlers.UpdateMetric)
	})

	r.Route("/value", func(r chi.Router) {
		r.Get("/{label:[a-zA-Z0-9]+}/{name:[a-zA-Z0-9]+}", handlers.ValueMetric)
	})

	return r
}
