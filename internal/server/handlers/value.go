package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
)

type ValueMetricHandler struct {
	config *config.Config
	repo   repositories.Repository
}

func (h *ValueMetricHandler) ValueMetricJSON(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var jm metric.JSONMetric
	if err := json.Unmarshal(data, &jm); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	i, err := metric.NewMetricIdentFromJSON(&jm)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := h.processMetric(res, i)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := m.Bytes()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if _, err = res.Write(body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *ValueMetricHandler) ValueMetricParams(res http.ResponseWriter, req *http.Request) {
	t := chi.URLParam(req, "type")
	n := chi.URLParam(req, "name")

	i, err := metric.NewMetricIdent(n, t)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := h.processMetric(res, i)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	if _, err = res.Write([]byte(m.String())); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *ValueMetricHandler) processMetric(res http.ResponseWriter, i metric.Ident) (metric.Metric, error) {
	m, err := h.repo.Get(i)
	if err != nil {
		if err == common.ErrMetricNotExists {
			http.Error(res, err.Error(), http.StatusNotFound)
		} else {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}
		return metric.Metric{}, err
	}

	return m, nil
}

func (h *ValueMetricHandler) GetHandlerJSON() HandlerFunc {
	return h.ValueMetricJSON
}

func (h *ValueMetricHandler) GetHandlerParams() HandlerFunc {
	return h.ValueMetricParams
}

func NewValueMetricHandler(c *config.Config, r repositories.Repository) *ValueMetricHandler {
	return &ValueMetricHandler{
		config: c,
		repo:   r,
	}
}
