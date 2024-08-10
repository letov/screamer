package collector

import (
	"log"
	"math/rand"
	"reflect"
	"runtime"
	"screamer/internal/collector/collector_maps"
	"screamer/internal/config"
	"screamer/internal/metric/kinds"
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

type MetricExport = map[string]string

const randomMetric = "RandomValue"
const pollCountMetric = "PollCount"

type MetricList interface {
	Update(n string, v interface{}) error
	Get(n string) (interface{}, error)
	Export() *MetricExport
}

type MetricLists struct {
	Counter MetricList
	Gauge   MetricList
}

var metricList MetricLists

func Init() {
	metricList = MetricLists{
		Counter: collector_maps.NewCounterMap(),
		Gauge:   collector_maps.NewGaugeMap(),
	}
}

func UpdateMetrics() {
	updateRuntimeMetrics()
	updateRandMetric()
	updatePollCountMetric()
}

func Export() map[kinds.Label]*MetricExport {
	return map[kinds.Label]*MetricExport{
		kinds.GaugeLabel:   metricList.Gauge.Export(),
		kinds.CounterLabel: metricList.Counter.Export(),
	}
}

func GetMetricNames() []string {
	return append(runtimeMetrics, randomMetric, pollCountMetric)
}

func updateRuntimeMetrics() {
	c := config.GetConfig()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	for _, fieldName := range runtimeMetrics {
		value := reflect.ValueOf(m)
		field := value.FieldByName(fieldName)
		v, err := toFloat64(field)
		if err == nil {
			_ = metricList.Gauge.Update(fieldName, v)
		} else if c.AgentLogEnable {
			log.Println("Cant parse metric", fieldName, err.Error())
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
	return 0, collector_maps.ErrKindExists
}

func updatePollCountMetric() {
	pc, _ := metricList.Counter.Get(pollCountMetric)
	var v int64
	if pc == nil {
		v = 0
	} else {
		v = pc.(int64) + 1
	}
	_ = metricList.Counter.Update(pollCountMetric, v)
}

func updateRandMetric() {
	_ = metricList.Gauge.Update(randomMetric, rand.Float64())
}
