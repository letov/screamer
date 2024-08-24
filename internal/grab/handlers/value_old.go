package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"screamer/internal/metric"
	"screamer/internal/storage"
	"screamer/internal/storage/repos/kinds"
)

func ValueMetricOld(res http.ResponseWriter, req *http.Request) {
	label := chi.URLParam(req, "label")
	name := chi.URLParam(req, "name")

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

	v, err := s.GetLastAsString(k, name)
	if err != nil {
		if err == kinds.ErrEmptyMetric {
			http.Error(res, err.Error(), http.StatusNotFound)
		} else {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	_, err = res.Write([]byte(v))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}
