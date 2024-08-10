package repos

import (
	"fmt"
	"screamer/internal/metric"
	"screamer/internal/storage/repos/mem_kinds"
)

type MetricStorage interface {
	Add(n string, v interface{}) error
	Get(n string) (interface{}, error)
	GetAsString(n string) (string, error)
	Debug() string
}

type MemStorage struct {
	Counter MetricStorage
	Gauge   MetricStorage
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Counter: mem_kinds.NewCounterStorage(),
		Gauge:   mem_kinds.NewGaugeStorage(),
	}
}

func (s *MemStorage) Add(m metric.Metric) error {
	switch m.Kind {
	case metric.Counter:
		return s.Counter.Add(m.Name, m.Value)
	case metric.Gauge:
		return s.Gauge.Add(m.Name, m.Value)
	}
	return mem_kinds.ErrUnknownMetricaIdent
}

func (s *MemStorage) Get(k metric.Kind, n string) (interface{}, error) {
	switch k {
	case metric.Counter:
		return s.Counter.Get(n)
	case metric.Gauge:
		return s.Gauge.Get(n)
	}
	return nil, mem_kinds.ErrUnknownMetricaIdent
}

func (s *MemStorage) GetAsString(k metric.Kind, n string) (string, error) {
	switch k {
	case metric.Counter:
		return s.Counter.GetAsString(n)
	case metric.Gauge:
		return s.Gauge.GetAsString(n)
	}
	return "", mem_kinds.ErrUnknownMetricaIdent
}

func (s *MemStorage) Debug() string {
	return fmt.Sprintf("StorageCounter: %v, StorageGauge: %v", s.Counter.Debug(), s.Gauge.Debug())
}
