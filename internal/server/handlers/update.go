package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"screamer/internal/common/metric"
	"screamer/internal/server/services"
	"time"
)

type UpdateMetricHandler struct {
	ms *services.MetricService
}

func (h *UpdateMetricHandler) Handler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := io.ReadAll(req.Body)
	_ = req.Body.Close()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var jm metric.JSONMetric
	err = json.Unmarshal(data, &jm)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := h.ms.UpdateMetricJSON(ctx, jm)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := m.Bytes()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.Header().Set("Content-Type", "application/json")
	if _, err = res.Write(body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewUpdateMetricHandler(ms *services.MetricService) *UpdateMetricHandler {
	return &UpdateMetricHandler{
		ms: ms,
	}
}
