package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/common/domain/metric"
	"screamer/internal/server/application/services"
	"time"
)

type ValueMetricHandler struct {
	ms *services.MetricService
}

func (h *ValueMetricHandler) Handler(res http.ResponseWriter, req *http.Request) {
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

	m, err := h.ms.ValueMetricJSON(ctx, jm)
	if err != nil {
		if errors.Is(err, common.ErrMetricNotExists) {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}
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

func NewValueMetricHandler(ms *services.MetricService) *ValueMetricHandler {
	return &ValueMetricHandler{
		ms: ms,
	}
}
