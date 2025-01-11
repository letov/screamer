package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"screamer/internal/common/application/dto"
	"screamer/internal/server/application/services"
	"time"
)

type UpdateBatchMetricHandler struct {
	ms *services.MetricService
}

func (h *UpdateBatchMetricHandler) Handler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := io.ReadAll(req.Body)
	_ = req.Body.Close()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var jms []dto.JSONMetric
	err = json.Unmarshal(data, &jms)
	if err != nil {
		return
	}

	err = h.ms.UpdateBatchMetricJSON(ctx, jms)
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

func NewUpdateBatchMetricHandler(ms *services.MetricService) *UpdateBatchMetricHandler {
	return &UpdateBatchMetricHandler{
		ms: ms,
	}
}
