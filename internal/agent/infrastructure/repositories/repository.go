package repositories

import (
	"context"
	"screamer/internal/common/domain/metric"
)

type Repository interface {
	Get(ctx context.Context, i metric.Ident) (metric.Metric, error)
	GetAll(ctx context.Context) []metric.Metric
	Update(ctx context.Context, m metric.Metric) (metric.Metric, error)
	Increase(ctx context.Context, i metric.Ident, v float64) (metric.Metric, error)
}
