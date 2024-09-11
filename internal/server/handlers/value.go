package handlers

import (
	"context"
	"io"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/server/services"
	"time"
)

type ValueMetricHandler struct {
	ms *services.MetricService
}

func (h *ValueMetricHandler) Handler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := h.ms.ValueMetricJSON(ctx, &data)
	if err != nil {
		if err == common.ErrMetricNotExists {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
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
