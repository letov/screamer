package metric_sources

import (
	"screamer/internal/common/metric"
)

const pollCount = "PollCount"

func getPollCountMetric() []*metric.Metric {
	pc := metric.NewCounter(pollCount, 0)
	return []*metric.Metric{pc}
}
