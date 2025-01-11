package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/common/application/dto"
	"screamer/internal/common/domain"
	"screamer/internal/server/application/services"
	"time"

	"github.com/go-chi/chi/v5"
)

type ValueMetricOldHandler struct {
	ms *services.MetricService
}

func (h *ValueMetricOldHandler) Handler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := chi.URLParam(req, "type")
	n := chi.URLParam(req, "name")

	m, err := domain.NewMetric(n, 0, domain.Type(t))

	jm, err := dto.NewJSONMetric(m)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	r, err := h.ms.ValueMetricJSON(ctx, jm)
	if err != nil {
		if errors.Is(err, common.ErrMetricNotExists) {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(r)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.Header().Set("Content-Type", "text/html")
	if _, err = res.Write(body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewValueMetricOldHandler(ms *services.MetricService) *ValueMetricOldHandler {
	return &ValueMetricOldHandler{
		ms: ms,
	}
}
