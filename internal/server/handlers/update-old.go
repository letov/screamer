package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"screamer/internal/server/services"
)

type UpdateMetricOldHandler struct {
	ms *services.MetricService
}

func (h *UpdateMetricOldHandler) Handler(res http.ResponseWriter, req *http.Request) {
	n := chi.URLParam(req, "name")
	v := chi.URLParam(req, "value")
	t := chi.URLParam(req, "type")

	body, err := h.ms.UpdateMetricParams(n, v, t)
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
