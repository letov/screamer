package repositories

import (
	"context"
	"screamer/internal/common/metric"
)

type Repository interface {
	Get(ctx context.Context, i metric.Ident) (metric.Metric, error)
	GetAll(ctx context.Context) []metric.Metric
	Add(ctx context.Context, m metric.Metric) (metric.Metric, error)
	Increase(ctx context.Context, m metric.Metric) (metric.Metric, error)
}
