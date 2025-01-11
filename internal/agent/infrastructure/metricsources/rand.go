package metricsources

import (
	"math/rand"
	"screamer/internal/common/domain/metric"
)

const randomMetric = "RandomValue"

func getRandMetric() []*metric.Metric {
	rm := metric.NewGauge(randomMetric, rand.Float64())
	return []*metric.Metric{rm}
}
