package store

import (
	"context"
	"errors"
	"math"
	"screamer/internal/common"
	"screamer/internal/common/domain"
	"sync"
)

type Memory struct {
	metrics map[domain.Ident][]domain.Metric
	sync.Mutex
}

func (mr *Memory) BatchUpdate(_ context.Context, _ []domain.Metric) error {
	//TODO implement me
	panic("implement me")
}

func (mr *Memory) GetAll(_ context.Context) []domain.Metric {
	res := make([]domain.Metric, 0)
	mr.Lock()
	for _, m := range mr.metrics {
		if len(m) > 0 {
			last := m[len(m)-1]
			res = append(res, last)
		}
	}
	mr.Unlock()
	return res
}

func (mr *Memory) Add(_ context.Context, m domain.Metric) (domain.Metric, error) {
	mr.Lock()
	mr.metrics[m.Ident] = append(mr.metrics[m.Ident], m)
	mr.Unlock()
	return m, nil
}

func (mr *Memory) Get(_ context.Context, i domain.Ident) (domain.Metric, error) {
	mv, ok := mr.metrics[i]
	if !ok {
		return domain.Metric{}, common.ErrMetricNotExists
	}
	if len(mv) == 0 {
		return domain.Metric{}, nil
	}
	return mv[len(mv)-1], nil
}

func (mr *Memory) Increase(ctx context.Context, m domain.Metric) (domain.Metric, error) {
	var _, frac float64
	_, frac = math.Modf(m.Value)
	if frac != 0 {
		return m, common.ErrInvalidValue
	}

	currentM, err := mr.Get(ctx, m.Ident)
	if err != nil && errors.Is(err, common.ErrMetricNotExists) {
		addM, e := domain.NewMetric(m.Ident.Name, m.Value, domain.Counter)
		if e != nil {
			return domain.Metric{}, e
		}
		return mr.Add(ctx, addM)
	}
	if err != nil {
		return currentM, err
	}
	m.Value += currentM.Value
	return mr.Add(ctx, m)
}

func NewMemory() *Memory {
	return &Memory{
		metrics: make(map[domain.Ident][]domain.Metric),
	}
}
