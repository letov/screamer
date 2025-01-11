package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"screamer/internal/common/application/dto"
	"screamer/internal/common/domain"
	"screamer/internal/server/application/services"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type UpdateMetricOldHandler struct {
	ms *services.MetricService
}

func (h *UpdateMetricOldHandler) Handler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	n := chi.URLParam(req, "name")
	v := chi.URLParam(req, "value")
	t := chi.URLParam(req, "type")

	vf, err := strconv.ParseFloat(v, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := domain.NewMetric(n, vf, domain.Type(t))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	jm, err := dto.NewJSONMetric(m)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	r, err := h.ms.UpdateMetricJSON(ctx, jm)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(r)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.Header().Set("Content-Type", "application/json")
	if _, err = res.Write(body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewUpdateMetricOldHandler(ms *services.MetricService) *UpdateMetricOldHandler {
	return &UpdateMetricOldHandler{
		ms: ms,
	}
}
