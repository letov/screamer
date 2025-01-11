package metricsources

import (
	"screamer/internal/common/domain/metric"
)

const pollCount = "PollCount"

func getPollCountMetric() []*metric.Metric {
	pc := metric.NewCounter(pollCount, 0)
	return []*metric.Metric{pc}
}
