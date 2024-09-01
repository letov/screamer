package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
	"screamer/internal/server/services"
	"strconv"
)

type UpdateMetricHandler struct {
	config        *config.Config
	repo          repositories.Repository
	backupService *services.BackupService
}

func (h *UpdateMetricHandler) UpdateMetricJSON(res http.ResponseWriter, req *http.Request) {
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

	m, err := metric.NewMetricFromJSON(&jm)
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

	body, err := newM.Bytes()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if _, err = res.Write(body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if h.config.Restore && h.config.StoreInterval == 0 {
		h.backupService.Save()
	}
}

func (h *UpdateMetricHandler) GetHandlerJSON() HandlerFunc {
	return h.UpdateMetricJSON
}

func (h *UpdateMetricHandler) GetHandlerParams() HandlerFunc {
	return h.UpdateMetricParams
}

func NewUpdateMetricHandler(c *config.Config, r repositories.Repository, bs *services.BackupService) *UpdateMetricHandler {
	return &UpdateMetricHandler{
		config:        c,
		repo:          r,
		backupService: bs,
	}
}
