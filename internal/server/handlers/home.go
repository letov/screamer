package handlers

import (
	"net/http"
	"screamer/internal/server/services"
)

type HomeHandler struct {
	ms *services.MetricService
}

func (h *HomeHandler) Handler(res http.ResponseWriter, _ *http.Request) {
	body := h.ms.Home()

	res.Header().Set("Content-Type", "text/html")
	_, err := res.Write(*body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewHomeHandler(ms *services.MetricService) *HomeHandler {
	return &HomeHandler{
		ms: ms,
	}
}
