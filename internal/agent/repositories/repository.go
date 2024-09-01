package repositories

import "screamer/internal/common/metric"

type Repository interface {
	Get(i metric.Ident) (metric.Metric, error)
	GetAll() []metric.Metric
	Update(m metric.Metric) (metric.Metric, error)
	Increase(i metric.Ident, v float64) (metric.Metric, error)
}
