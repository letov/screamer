package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"screamer/internal/metric"
	"screamer/internal/storage"
)

func UpdateMetric(res http.ResponseWriter, req *http.Request) {
	label := chi.URLParam(req, "label")
	name := chi.URLParam(req, "name")

	m, err := metric.NewMetric(metric.Raw{
		Label: label,
		Name:  name,
		Value: chi.URLParam(req, "value"),
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
	err = s.Add(m)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}
