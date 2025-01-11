package metricsources

import (
	"screamer/internal/common/domain"
)

const pollCount = "PollCount"

func getPollCountMetric() []domain.Metric {
	pc, _ := domain.NewMetric(pollCount, 0, domain.Counter)
	return []domain.Metric{pc}
}
