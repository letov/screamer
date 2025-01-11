package repo

import (
	"context"
	"screamer/internal/common/domain"
)

type Repository interface {
	Get(ctx context.Context, i domain.Ident) (domain.Metric, error)
	GetAll(ctx context.Context) []domain.Metric
	Update(ctx context.Context, m domain.Metric) (domain.Metric, error)
	Increase(ctx context.Context, i domain.Ident, v float64) (domain.Metric, error)
}
