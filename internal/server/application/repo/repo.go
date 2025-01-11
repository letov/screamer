package repo

import (
	"context"
	"screamer/internal/common/domain"
)

type Repository interface {
	Get(ctx context.Context, i domain.Ident) (domain.Metric, error)
	GetAll(ctx context.Context) []domain.Metric
	Add(ctx context.Context, m domain.Metric) (domain.Metric, error)
	Increase(ctx context.Context, m domain.Metric) (domain.Metric, error)
	BatchUpdate(ctx context.Context, m []domain.Metric) error
}
