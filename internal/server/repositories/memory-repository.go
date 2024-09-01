package repositories

import (
	"math"
	"screamer/internal/common"
	"screamer/internal/common/metric"
	"sync"
)

type MemoryRepository struct {
	metrics map[metric.Ident][]metric.Metric
	sync.Mutex
}

func (mr *MemoryRepository) GetAll() []metric.Metric {
	res := make([]metric.Metric, 0)
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

func (mr *MemoryRepository) Add(m metric.Metric) (metric.Metric, error) {
	mr.Lock()
	mr.metrics[m.Ident] = append(mr.metrics[m.Ident], m)
	mr.Unlock()
	return m, nil
}

func (mr *MemoryRepository) Get(i metric.Ident) (metric.Metric, error) {
	mv, ok := mr.metrics[i]
	if !ok {
		return metric.Metric{}, common.ErrMetricNotExists
	}
	if len(mv) == 0 {
		return metric.Metric{}, nil
	}
	return mv[len(mv)-1], nil
}

func (mr *MemoryRepository) Increase(m metric.Metric) (metric.Metric, error) {
	var _, frac float64
	_, frac = math.Modf(m.Value)
	if frac != 0 {
		return m, common.ErrInvalidValue
	}

	currentM, err := mr.Get(m.Ident)
	if err != nil && err == common.ErrMetricNotExists {
		return mr.Add(*metric.NewCounter(m.Ident.Name, m.Value))
	}
	if err != nil {
		return currentM, err
	}
	m.Value += currentM.Value
	return mr.Add(m)
}

func NewMemoryRepository() Repository {
	return &MemoryRepository{
		metrics: make(map[metric.Ident][]metric.Metric),
	}
}
