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
	"strconv"
)

type ValueMetricHandler struct {
	config *config.Config
	repo   repositories.Repository
}

func (h *ValueMetricHandler) ValueMetricJson(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var jm metric.JsonMetric
	if err := json.Unmarshal(data, &jm); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	i, err := metric.NewMetricIdentFromJson(&jm)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := h.processMetric(res, i)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	newJ, err := m.Json()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(newJ)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(body)
	if err != nil {
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

	vs := strconv.FormatFloat(m.Value, 'f', 6, 64)

	res.Header().Set("Content-Type", "text/html")
	_, err = res.Write([]byte(vs))
	if err != nil {
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

func (h *ValueMetricHandler) GetHandlerJson() HandlerFunc {
	return h.ValueMetricJson
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