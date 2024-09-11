package handlers

import (
	"context"
	"io"
	"net/http"
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
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := h.ms.UpdateMetricJSON(ctx, &data)
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
