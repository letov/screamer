package mem_kinds

import (
	"fmt"
	"time"
)

type Gauge struct {
	Timestamp int64
	Value     float64
}

type GaugeStorageType = map[string][]Gauge

type GaugeStorage struct {
	Storage GaugeStorageType
}

func NewGaugeStorage() *GaugeStorage {
	return &GaugeStorage{
		Storage: make(GaugeStorageType),
	}
}

func (s *GaugeStorage) Add(n string, v interface{}) error {
	data, ok := v.(float64)
	if !ok {
		return ErrInvalidDataType
	}
	s.Storage[n] = append(s.Storage[n], Gauge{
		Timestamp: time.Now().Unix(),
		Value:     data,
	})
	return nil
}

func (s *GaugeStorage) Get(n string) (interface{}, error) {
	if _, ok := s.Storage[n]; !ok {
		return nil, ErrEmptyMetric
	}
	l := len(s.Storage[n])
	if l == 0 {
		return nil, ErrEmptyMetric
	}
	return s.Storage[n][l-1].Value, nil
}

func (s *GaugeStorage) GetAsString(n string) (string, error) {
	v, err := s.Get(n)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v), nil
}

func (s *GaugeStorage) GetAllLastAsString() (*map[string]string, error) {
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
