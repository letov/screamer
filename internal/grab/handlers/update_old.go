package handlers

import (
	"net/http"
)

func UpdateMetricOld(res http.ResponseWriter, req *http.Request) {
	//label := chi.URLParam(req, "label")
	//name := chi.URLParam(req, "name")
	//value := chi.URLParam(req, "value")
	//
	//m, err := metric.NewMetric(metric.Raw{
	//	Label: label,
	//	Name:  name,
	//	Value: value,
	//})
	//
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//s := storage.GetStorage()
	//if s == nil {
	//	http.Error(res, ErrNoStorage.Error(), http.StatusBadRequest)
	//	return
	//}
	//_, err = s.Add(m)
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//_, err = res.Write([]byte(""))
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
}
