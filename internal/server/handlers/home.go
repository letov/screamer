package handlers

import (
	"context"
	"net/http"
	"screamer/internal/server/services"
	"time"
)

type HomeHandler struct {
	ms *services.MetricService
}

func (h *HomeHandler) Handler(res http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := h.ms.Home(ctx)

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
