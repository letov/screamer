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
	_ = storage.GetStorage().Add(m)
}
