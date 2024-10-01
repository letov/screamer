package metricsources

import "screamer/internal/common/metric"

type MetricSource = func() []*metric.Metric

func NewMetricSources() []MetricSource {
	return []MetricSource{
		getRuntimeMetrics,
		getPollCountMetric,
		getRandMetric,
		getGopsutilMetric,
	}
}
