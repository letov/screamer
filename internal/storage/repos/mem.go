package repos

import (
	"screamer/internal/metric"
	"screamer/internal/storage/repos/kinds"
)

type MetricStorage interface {
	Add(n string, v interface{}) (interface{}, error)
	Increase(n string, v interface{}) (interface{}, error)
	GetLast(n string) (interface{}, error)
	GetLastAsString(n string) (string, error)
	GetAllLastAsString() (*kinds.MetricList, error)
}

type MemStorage struct {
	Counter MetricStorage
	Gauge   MetricStorage
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Counter: kinds.NewCounterStorage(),
		Gauge:   kinds.NewGaugeStorage(),
	}
}

func (s *MemStorage) Add(m metric.Metric) (interface{}, error) {
	switch m.Kind {
	case metric.Counter:
		return s.Counter.Increase(m.Name, m.Value)
	case metric.Gauge:
		return s.Gauge.Add(m.Name, m.Value)
	}
	return nil, kinds.ErrUnknownMetricaIdent
}

func (s *MemStorage) GetLast(k metric.Kind, n string) (interface{}, error) {
	switch k {
	case metric.Counter:
		return s.Counter.GetLast(n)
	case metric.Gauge:
		return s.Gauge.GetLast(n)
	}
	return nil, kinds.ErrUnknownMetricaIdent
}

func (s *MemStorage) GetLastAsString(k metric.Kind, n string) (string, error) {
	switch k {
	case metric.Counter:
		return s.Counter.GetLastAsString(n)
	case metric.Gauge:
		return s.Gauge.GetLastAsString(n)
	}
	return "", kinds.ErrUnknownMetricaIdent
}

func (s *MemStorage) GetAllLastAsString(k metric.Kind) (*kinds.MetricList, error) {
	switch k {
	case metric.Counter:
		return s.Counter.GetAllLastAsString()
	case metric.Gauge:
		return s.Gauge.GetAllLastAsString()
	}
	return nil, kinds.ErrUnknownMetricaIdent
}
