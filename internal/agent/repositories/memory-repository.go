package repositories

import (
	"context"
	"screamer/internal/common"
	"screamer/internal/common/metric"
	"sync"
)

type MemoryRepository struct {
	metrics map[metric.Ident]metric.Metric
	sync.Mutex
}

func (mr *MemoryRepository) GetAll(_ context.Context) []metric.Metric {
	res := make([]metric.Metric, 0)
	mr.Lock()
	for _, m := range mr.metrics {
		res = append(res, m)
	}
	mr.Unlock()
	return res
}

func (mr *MemoryRepository) Update(_ context.Context, m metric.Metric) (metric.Metric, error) {
	mr.Lock()
	mr.metrics[m.Ident] = m
	mr.Unlock()
	return m, nil
}

func (mr *MemoryRepository) Get(_ context.Context, i metric.Ident) (metric.Metric, error) {
	mv, ok := mr.metrics[i]
	if !ok {
		return metric.Metric{}, common.ErrMetricNotExists
	}
	return mv, nil
}

func (mr *MemoryRepository) Increase(ctx context.Context, i metric.Ident, v float64) (metric.Metric, error) {
	mv, err := mr.Get(ctx, i)
	if err != nil && err == common.ErrMetricNotExists {
		_, err = mr.Update(ctx, *metric.NewCounter(i.Name, v))
		return mv, err
	}
	if err != nil {
		return mv, err
	}
	mv.Value += v
	return mr.Update(ctx, mv)
}

func NewMemoryRepository() Repository {
	return &MemoryRepository{
		metrics: make(map[metric.Ident]metric.Metric),
	}
}
