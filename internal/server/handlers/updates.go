package handlers

import (
	"context"
	"io"
	"net/http"
	"screamer/internal/server/services"
	"time"
)

type UpdatesMetricHandler struct {
	ms *services.MetricService
}

func (h *UpdatesMetricHandler) Handler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ms.UpdatesMetricJSON(ctx, &data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	body := []byte("OK")
	if _, err = res.Write(body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewUpdatesMetricHandler(ms *services.MetricService) *UpdatesMetricHandler {
	return &UpdatesMetricHandler{
		ms: ms,
	}
}
