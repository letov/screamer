package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"screamer/internal/collector"
	"screamer/internal/metric"
	"screamer/internal/storage"
	"slices"
)

func ValueMetric(res http.ResponseWriter, req *http.Request) {
	label := chi.URLParam(req, "label")
	name := chi.URLParam(req, "name")

	if !isValidMetricName(name) {
		http.Error(res, ErrNoMetric.Error(), http.StatusNotFound)
		return
	}

	k, err := metric.LabelToKind(label)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	s := storage.GetStorage()
	if s == nil {
		http.Error(res, ErrNoStorage.Error(), http.StatusBadRequest)
		return
	}
	v, err := s.GetAsString(k, name)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = res.Write([]byte(v))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func isValidMetricName(n string) bool {
	ns := collector.GetMetricNames()
	return slices.Contains(ns, n)
}
