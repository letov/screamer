package metric_sources

import (
	"math/rand"
	"screamer/internal/common/metric"
)

const randomMetric = "RandomValue"

func gewRandMetric() []*metric.Metric {
	rm := metric.NewGauge(randomMetric, rand.Float64())
	return []*metric.Metric{rm}
}
