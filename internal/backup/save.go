package backup

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"screamer/internal/metric"
	"screamer/internal/server/config"
	"screamer/internal/storage"
	"screamer/internal/storage/repos/kinds"
)

type JsonBUMetric struct {
	Kind  metric.Kind `json:"kind"`
	Name  string      `json:"name"`
	Value string      `json:"value"`
}

type JsonBUMetricList struct {
	Array []JsonBUMetric
}

var ks = []metric.Kind{
	metric.Counter,
	metric.Gauge,
}

func Save() {
	s := storage.GetStorage()
	c := config.GetConfig()
	for _, k := range ks {
		ml, err := s.GetAllLastAsString(k)
		if err == nil {
			err = toFile(ml, k, *c.FileStoragePath)
		}
		if err != nil && *c.ServerLogEnable {
			log.Println("Save backup error:", err.Error())
		}
	}
}

// TODO: need refactor
//func Load() {
//s := storage.GetStorage()
//c := config.GetConfigS()
//for _, k := range ks {
//	j, err := fromFile(k, c.FileStoragePath)
//	if err != nil && c.ServerLogEnable {
//		log.Println("Load backup error:", err.Error())
//		continue
//	}
//	ml := jsonToMetricList(j)
//	for n ,v := range ml {
//		m, err := metric.NewMetric(metric.Raw{
//			Label: KindToLabel,
//			Name:  jm.ID,
//			Value: value,
//		})
//		if err != nil && c.ServerLogEnable {
//			log.Println("Load backup error:", err.Error())
//			continue
//		}
//		s.Add()
//	}
//ml, err := s.GetAllLastAsString(k)
//if err == nil {
//	err = toFile(ml, k, c.FileStoragePath)
//}
//if err != nil && c.ServerLogEnable {
//	log.Println("Save backup error:", err.Error())
//}
//
//	}
//}

func toFile(ml *kinds.MetricList, k metric.Kind, fileStoragePath string) error {
	j := metricListToJson(k, ml)
	data, err := json.MarshalIndent(j, "", "   ")
	if err != nil {
		return err
	}
	fp, err := getFilePath(k, fileStoragePath)
	if err != nil {
		return err
	}
	return os.WriteFile(fp, data, 0666)
}

func fromFile(k metric.Kind, fileStoragePath string) (j *JsonBUMetricList, err error) {
	fp, err := getFilePath(k, fileStoragePath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, j)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func getFilePath(k metric.Kind, fileStoragePath string) (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v/%v/%v", cur, fileStoragePath, k), nil
}

func metricListToJson(k metric.Kind, ml *kinds.MetricList) *JsonBUMetricList {
	var arr []JsonBUMetric

	for n, v := range *ml {
		arr = append(arr, JsonBUMetric{
			Kind:  k,
			Name:  n,
			Value: v,
		})
	}

	return &JsonBUMetricList{
		Array: arr,
	}
}

func jsonToMetricList(j *JsonBUMetricList) *kinds.MetricList {
	var ml kinds.MetricList

	for _, jm := range (*j).Array {
		ml[jm.Name] = jm.Value
	}

	return &ml
}
