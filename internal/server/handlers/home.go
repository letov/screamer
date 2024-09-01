package handlers

import (
	"fmt"
	"net/http"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
)

type HomeHandler struct {
	config *config.Config
	repo   repositories.Repository
}

func (h *HomeHandler) Home(res http.ResponseWriter, _ *http.Request) {
	r := "<html><body>"
	r += "<h1>Metrics</h1>"
	ms := h.repo.GetAll()
	for _, m := range ms {
		r += fmt.Sprintf("<p>%v: %v</p>", m.Ident.Name, m.Value)
	}
	r += "</body></html>"

	res.Header().Set("Content-Type", "text/html")
	_, err := res.Write([]byte(r))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *HomeHandler) GetHandler() HandlerFunc {
	return h.Home
}

func NewHomeHandler(c *config.Config, r repositories.Repository) *HomeHandler {
	return &HomeHandler{
		config: c,
		repo:   r,
	}
}
