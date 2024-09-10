package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/server/services"
)

type ValueMetricOldHandler struct {
	ms *services.MetricService
}

func (h *ValueMetricOldHandler) Handler(res http.ResponseWriter, req *http.Request) {
	t := chi.URLParam(req, "type")
	n := chi.URLParam(req, "name")

	body, err := h.ms.ValueMetricParams(n, t)
	if err != nil {
		if err == common.ErrMetricNotExists {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	if _, err = res.Write(*body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewValueMetricOldHandler(ms *services.MetricService) *ValueMetricOldHandler {
	return &ValueMetricOldHandler{
		ms: ms,
	}
}
