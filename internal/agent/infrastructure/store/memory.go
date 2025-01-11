package store

import (
	"context"
	"errors"
	"screamer/internal/agent/application/repo"
	"screamer/internal/common"
	"screamer/internal/common/domain"
	"sync"
)

type MemoryRepository struct {
	metrics map[domain.Ident]domain.Metric
	sync.Mutex
}

func (mr *MemoryRepository) GetAll(_ context.Context) []domain.Metric {
	res := make([]domain.Metric, 0)
	mr.Lock()
	for _, m := range mr.metrics {
		res = append(res, m)
	}
	mr.Unlock()
	return res
}

func (mr *MemoryRepository) Update(_ context.Context, m domain.Metric) (domain.Metric, error) {
	mr.Lock()
	mr.metrics[m.Ident] = m
	mr.Unlock()
	return m, nil
}

func (mr *MemoryRepository) Get(_ context.Context, i domain.Ident) (domain.Metric, error) {
	mv, ok := mr.metrics[i]
	if !ok {
		return domain.Metric{}, common.ErrMetricNotExists
	}
	return mv, nil
}

func (mr *MemoryRepository) Increase(ctx context.Context, i domain.Ident, v float64) (domain.Metric, error) {
	mv, err := mr.Get(ctx, i)
	if err != nil && errors.Is(err, common.ErrMetricNotExists) {
		addM, e := domain.NewMetric(i.Name, v, domain.Counter)
		if e != nil {
			return domain.Metric{}, e
		}
		_, err = mr.Update(ctx, addM)
		return mv, err
	}
	if err != nil {
		return mv, err
	}
	mv.Value += v
	return mr.Update(ctx, mv)
}

func NewMemoryRepository() repo.Repository {
	return &MemoryRepository{
		metrics: make(map[domain.Ident]domain.Metric),
	}
}
