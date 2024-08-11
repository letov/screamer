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

func (s *CounterStorage) Add(n string, v interface{}) (interface{}, error) {
	data, ok := v.(int64)
	if !ok {
		return nil, ErrInvalidDataType
	}
	s.Storage[n] = append(s.Storage[n], Counter{
		Timestamp: time.Now().Unix(),
		Value:     data,
	})
	return data, nil
}

func (s *CounterStorage) Increase(n string, v interface{}) (interface{}, error) {
	data, ok := v.(int64)
	if !ok {
		return nil, ErrInvalidDataType
	}
	cur, err := s.GetLast(n)
	if err == ErrEmptyMetric {
		cur = 0
	} else {
		return nil, err
	}
	incr := cur.(int64) + data
	return s.Add(n, incr)
}

func (s *CounterStorage) GetLast(n string) (interface{}, error) {
	if _, ok := s.Storage[n]; !ok {
		return nil, ErrEmptyMetric
	}
	l := len(s.Storage[n])
	if l == 0 {
		return nil, ErrEmptyMetric
	}
	return s.Storage[n][l-1].Value, nil
}

func (s *CounterStorage) GetLastAsString(n string) (string, error) {
	v, err := s.GetLast(n)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v), nil
}

func (s *CounterStorage) GetAllLastAsString() (*map[string]string, error) {
	res := make(map[string]string)

	for n, vs := range s.Storage {
		l := len(vs)
		if len(vs) > 0 {
			v := vs[l-1].Value
			res[n] = fmt.Sprintf("%v", v)
		}
	}

	return &res, nil
}
