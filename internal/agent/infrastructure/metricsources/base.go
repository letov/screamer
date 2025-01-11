package metricsources

import (
	"screamer/internal/common/domain"
)

type MetricSource = func() []domain.Metric

func NewMetricSources() []MetricSource {
	return []MetricSource{
		getRuntimeMetrics,
		getPollCountMetric,
		getRandMetric,
		getGopsutilMetric,
	}
}
