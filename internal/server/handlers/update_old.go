package handlers

import (
	"context"
	"net/http"
	"screamer/internal/server/services"
	"time"

	"github.com/go-chi/chi/v5"
)

type UpdateMetricOldHandler struct {
	ms *services.MetricService
}

func (h *UpdateMetricOldHandler) Handler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	n := chi.URLParam(req, "name")
	v := chi.URLParam(req, "value")
	t := chi.URLParam(req, "type")

	body, err := h.ms.UpdateMetricParams(ctx, n, v, t)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if _, err = res.Write(*body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewUpdateMetricOldHandler(ms *services.MetricService) *UpdateMetricOldHandler {
	return &UpdateMetricOldHandler{
		ms: ms,
	}
}
