package handlers

import (
	"net/http"
)

func UpdateMetric(res http.ResponseWriter, req *http.Request) {
	//data, err := io.ReadAll(req.Body)
	//
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//var jm JsonMetric
	//if err := json.Unmarshal(data, &jm); err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//var value string
	//if jm.Value != nil {
	//	value = fmt.Sprintf("%v", *jm.Value)
	//} else {
	//	value = fmt.Sprintf("%v", *jm.Delta)
	//}
	//
	//m, err := metric.NewMetric(metric.Raw{
	//	Label: jm.MType,
	//	Name:  jm.ID,
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
	//newV, err := s.Add(m)
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//body, err := GetMarshal(newV, &jm)
	//
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//res.Header().Set("Content-Type", "application/json")
	//res.WriteHeader(http.StatusOK)
	//_, err = res.Write(body)
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusBadRequest)
	//	return
	//}
}
