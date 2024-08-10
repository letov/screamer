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

type MetricExport = map[string]string

type MetricMap interface {
	Update(n string, v interface{}) error
	Get(n string) (interface{}, error)
	Export() MetricExport
}

type MetricMaps struct {
	Counter MetricMap
	Gauge   MetricMap
}

var metricMap MetricMaps

func Init() {
	metricMap = MetricMaps{
		Counter: collector_maps.NewCounterMap(),
		Gauge:   collector_maps.NewGaugeMap(),
	}
}

func UpdateMetrics() {
	updateRuntimeMetrics()
	_ = metricMap.Gauge.Update(RandomMetric, rand.Float64())
	increaseCountMetric(PollCountMetric)
}

func Export() map[kinds.Label]MetricExport {
	return map[kinds.Label]MetricExport{
		kinds.GaugeLabel:   metricMap.Gauge.Export(),
		kinds.CounterLabel: metricMap.Counter.Export(),
	}
}

func updateRuntimeMetrics() {
	c := config.GetConfig()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	for _, fieldName := range RuntimeMetrics {
		value := reflect.ValueOf(m)
		field := value.FieldByName(fieldName)
		v, err := toFloat64(field)
		if err == nil {
			_ = metricMap.Gauge.Update(fieldName, v)
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

func increaseCountMetric(n string) {
	m, _ := metricMap.Counter.Get(n)
	var v int64
	if m == nil {
		v = 0
	} else {
		v = m.(int64) + 1
	}
	_ = metricMap.Counter.Update(n, v)
}
