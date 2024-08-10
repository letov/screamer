package mem_kinds

import (
	"fmt"
	"time"
)

type Counter struct {
	Timestamp int64
	Value     int64
}

type CounterStorageType = map[string][]Counter

type CounterStorage struct {
	Storage CounterStorageType
}

func NewCounterStorage() *CounterStorage {
	return &CounterStorage{
		Storage: make(CounterStorageType),
	}
}

func (s *CounterStorage) Add(n string, v interface{}) error {
	data, ok := v.(int64)
	if !ok {
		return ErrInvalidDataType
	}
	s.Storage[n] = append(s.Storage[n], Counter{
		Timestamp: time.Now().Unix(),
		Value:     data,
	})
	return nil
}

func (s *CounterStorage) Get(n string) (interface{}, error) {
	if _, ok := s.Storage[n]; !ok {
		return nil, ErrEmptyMetric
	}
	l := len(s.Storage[n])
	if l == 0 {
		return nil, ErrEmptyMetric
	}
	return s.Storage[n][l-1].Value, nil
}

func (s *CounterStorage) GetAsString(n string) (string, error) {
	v, err := s.Get(n)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v), nil
}

func (s *CounterStorage) Debug() string {
	return fmt.Sprintf("%v", s.Storage)
}
