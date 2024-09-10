package handlers

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/server/services"
)

type ValueMetricHandler struct {
	ms *services.MetricService
}

func (h *ValueMetricHandler) ValueMetricJSON(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := h.ms.ValueMetricJSON(&data)
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

func (h *ValueMetricHandler) ValueMetricParams(res http.ResponseWriter, req *http.Request) {
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

func NewValueMetricHandler(ms *services.MetricService) *ValueMetricHandler {
	return &ValueMetricHandler{
		ms: ms,
	}
}
