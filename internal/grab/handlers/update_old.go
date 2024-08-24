package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"screamer/internal/metric"
	"screamer/internal/storage"
)

func UpdateMetricOld(res http.ResponseWriter, req *http.Request) {
	label := chi.URLParam(req, "label")
	name := chi.URLParam(req, "name")
	value := chi.URLParam(req, "value")

	m, err := metric.NewMetric(metric.Raw{
		Label: label,
		Name:  name,
		Value: value,
	})

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	s := storage.GetStorage()
	if s == nil {
		http.Error(res, ErrNoStorage.Error(), http.StatusBadRequest)
		return
	}

	_, err = s.Add(m)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = res.Write([]byte("OK"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}
