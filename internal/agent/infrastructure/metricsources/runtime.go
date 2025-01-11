package metricsources

import (
	"reflect"
	"runtime"
	"screamer/internal/common"
	"screamer/internal/common/domain"
)

func getRuntimeMetrics() []domain.Metric {
	metrics := make([]domain.Metric, 0)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	for _, n := range *getRuntimeMetricNames() {
		value := reflect.ValueOf(m)
		field := value.FieldByName(n)
		if v, err := toFloat64(field); err == nil {
			m, _ := domain.NewMetric(n, v, domain.Gauge)
			metrics = append(metrics, m)
		}
	}
	return metrics
}

func getRuntimeMetricNames() *[]string {
	return &[]string{
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
}

func toFloat64(field reflect.Value) (float64, error) {
	switch field.Kind() {
	case reflect.Float64:
		return field.Float(), nil
	case reflect.Uint32:
		return float64(field.Uint()), nil
	case reflect.Uint64:
		return float64(field.Uint()), nil
	default:
		return 0, common.ErrTypeNotExists
	}
}
