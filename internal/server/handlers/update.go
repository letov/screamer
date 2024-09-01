package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
	"strconv"
)

type UpdateMetricHandler struct {
	config *config.Config
	repo   repositories.Repository
}

func (h *UpdateMetricHandler) UpdateMetricJson(res http.ResponseWriter, req *http.Request) {
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

	m, err := metric.NewMetricFromJson(&jm)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	h.processMetric(res, m)
}

func (h *UpdateMetricHandler) UpdateMetricParams(res http.ResponseWriter, req *http.Request) {
	n := chi.URLParam(req, "name")
	vs := chi.URLParam(req, "value")
	t := chi.URLParam(req, "type")

	v, err := strconv.ParseFloat(vs, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := metric.NewMetric(n, v, t)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	h.processMetric(res, m)
}

func (h *UpdateMetricHandler) processMetric(res http.ResponseWriter, m *metric.Metric) {
	var newM metric.Metric
	var err error

	if m.Ident.Type == metric.Counter {
		newM, err = h.repo.Increase(*m)
	} else {
		newM, err = h.repo.Add(*m)
	}
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	newJm, err := newM.Json()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(newJm)
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

	//c := config.GetConfig()
	//if *c.Restore && *c.StoreInterval == 0 {
	//	backup.Save()
	//}
}

func (h *UpdateMetricHandler) GetHandlerJson() HandlerFunc {
	return h.UpdateMetricJson
}

func (h *UpdateMetricHandler) GetHandlerParams() HandlerFunc {
	return h.UpdateMetricParams
}

func NewUpdateMetricHandler(c *config.Config, r repositories.Repository) *UpdateMetricHandler {
	return &UpdateMetricHandler{
		config: c,
		repo:   r,
	}
}
