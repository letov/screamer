package collector

import (
	"errors"
	"log"
	"math/rand"
	"reflect"
	"runtime"
)

var runtimeMetrics = []string{
	"Alloc",
	"BuckHashSys",
	"Frees",
	"GCCPUFraction",
	"GCSys",
	"HeapAlloc",
	"HeapIdle",
	"HeapInuse",
	"HeapObjects",
	"HeapReleased",
	"HeapSys",
	"LastGC",
	"Lookups",
	"MCacheInuse",
	"MCacheSys",
	"MSpanInuse",
	"MSpanSys",
	"Mallocs",
	"NextGC",
	"NumForcedGC",
	"NumGC",
	"OtherSys",
	"PauseTotalNs",
	"StackInuse",
	"StackSys",
	"Sys",
	"TotalAlloc",
}

type Gauge = map[string]float64
type Counter = map[string]int64

var gauge Gauge
var counter Counter

func Init() {
	gauge = make(map[string]float64)
	counter = make(map[string]int64)
}

func UpdateMetrics() {
	updateRuntimeMetrics()
	gauge["RandomValue"] = rand.Float64()
	counter["PollCount"]++
}

func GetMetrics() (*Gauge, *Counter) {
	return &gauge, &counter
}

func updateRuntimeMetrics() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	for _, fieldName := range runtimeMetrics {
		value := reflect.ValueOf(m)
		field := value.FieldByName(fieldName)
		v, err := toFloat64(field)
		if err == nil {
			gauge[fieldName] = v
		} else {
			log.Fatal(err.Error())
		}
	}
}

func toFloat64(field reflect.Value) (float64, error) {
	switch field.Kind() {
	case reflect.Float64:
		return field.Float(), nil
	case reflect.Uint32:
		return float64(field.Uint()), nil
	case reflect.Uint64:
		return float64(field.Uint()), nil
	}
	return 0, errors.New("unknown metric kind")
}
