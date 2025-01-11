package metricsources

import (
	"math/rand"
	"screamer/internal/common/domain"
)

const randomMetric = "RandomValue"

func getRandMetric() []domain.Metric {
	rm, _ := domain.NewMetric(randomMetric, rand.Float64(), domain.Gauge)
	return []domain.Metric{rm}
}
