package repos

import (
	"fmt"
	"screamer/internal/metric"
	"time"
)

type MemCounter struct {
	Timestamp int64
	Name      string
	Value     int64
}

type MemGauge struct {
	Timestamp int64
	Name      string
	Value     float64
}

type MemStorage struct {
	StorageCounter []MemCounter
	StorageGauge   []MemGauge
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		StorageCounter: []MemCounter{},
		StorageGauge:   []MemGauge{},
	}
}

func (s *MemStorage) Add(m metric.Metric) error {
	switch m.Ident {
	case metric.CounterIdent:
		data, ok := m.Value.(int64)
		if !ok {
			return ErrInvalidDataType
		}
		s.StorageCounter = append(s.StorageCounter, MemCounter{
			Timestamp: time.Now().Unix(),
			Name:      m.Name,
			Value:     data,
		})
		return nil
	case metric.GaugeIdent:
		data, ok := m.Value.(float64)
		if !ok {
			return ErrInvalidDataType
		}
		s.StorageGauge = append(s.StorageGauge, MemGauge{
			Timestamp: time.Now().Unix(),
			Name:      m.Name,
			Value:     data,
		})
		return nil
	}
	return ErrUnknownMetricaIdent
}

func (s *MemStorage) Debug() string {
	return fmt.Sprintf("StorageCounter: %v, StorageGauge: %v", s.StorageCounter, s.StorageGauge)
}
