package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"screamer/internal/metric"
	"screamer/internal/metric/validators"
	"screamer/internal/storage"
)

const UpdateRoute = `/update/{label:[a-zA-Z]+}/{name:[a-zA-Z]+}/{value}`

func UpdateHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "", http.StatusNotFound)
		return
	}
	params := mux.Vars(req)
	m, err := metric.NewMetric(metric.Raw{
		Label: params["label"],
		Name:  params["name"],
		Value: params["value"],
	})

	switch err {
	case validators.ErrUnknownMetricType:
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	case validators.ErrIncorrectMetricValue:
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	s := storage.GetStorage()
	if s == nil {
		http.Error(res, "there is no init storage", http.StatusBadRequest)
		return
	}
	err = s.Add(m)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}
