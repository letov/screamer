package handlers

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"screamer/internal/server/services"
)

type UpdateMetricHandler struct {
	ms *services.MetricService
}

func (h *UpdateMetricHandler) UpdateMetricJSON(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := h.ms.UpdateMetricJSON(&data)
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

func (h *UpdateMetricHandler) UpdateMetricParams(res http.ResponseWriter, req *http.Request) {
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

func NewUpdateMetricHandler(ms *services.MetricService) *UpdateMetricHandler {
	return &UpdateMetricHandler{
		ms: ms,
	}
}
