package repositories

import "screamer/internal/common/metric"

type Repository interface {
	Get(i metric.Ident) (metric.Metric, error)
	GetAll() []metric.Metric
	Add(m metric.Metric) (metric.Metric, error)
	Increase(m metric.Metric) (metric.Metric, error)
}
