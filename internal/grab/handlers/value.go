package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"screamer/internal/metric"
	"screamer/internal/storage"
	"screamer/internal/storage/repos/kinds"
)

func ValueMetric(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var jm JsonMetric
	if err := json.Unmarshal(data, &jm); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	k, err := metric.LabelToKind(jm.MType)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	s := storage.GetStorage()
	if s == nil {
		http.Error(res, ErrNoStorage.Error(), http.StatusBadRequest)
		return
	}
	v, err := s.GetLast(k, jm.ID)
	if err != nil {
		if err == kinds.ErrEmptyMetric {
			http.Error(res, err.Error(), http.StatusNotFound)
		} else {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}
		return
	}

	body, err := GetMarshal(v, &jm)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = res.Write(body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}
