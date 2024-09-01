package handlers

import "net/http"

type HandlerFunc = func(res http.ResponseWriter, req *http.Request)

//
//import (
//	"encoding/json"
//	"screamer/internal/old/collector/maps"
//)
//
//type JsonMetric struct {
//	ID    string   `json:"id"`
//	MType string   `json:"type"`
//	Delta *int64   `json:"delta,omitempty"`
//	Value *float64 `json:"value,omitempty"`
//}
//
//func GetMarshal(v interface{}, jm *JsonMetric) ([]byte, error) {
//	var body []byte
//	var err error = nil
//	switch val := v.(type) {
//	case int64:
//		body, err = json.Marshal(JsonMetric{
//			ID:    jm.ID,
//			MType: jm.MType,
//			Delta: &val,
//		})
//		break
//	case float64:
//		body, err = json.Marshal(JsonMetric{
//			ID:    jm.ID,
//			MType: jm.MType,
//			Value: &val,
//		})
//		break
//	default:
//		err = maps.ErrTypecast
//		break
//	}
//	return body, err
//}
