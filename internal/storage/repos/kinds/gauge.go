package kinds

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

func (s *GaugeStorage) Add(n string, v interface{}) (interface{}, error) {
	data, ok := v.(float64)
	if !ok {
		return nil, ErrInvalidDataType
	}
	s.Storage[n] = append(s.Storage[n], Gauge{
		Timestamp: time.Now().Unix(),
		Value:     data,
	})
	return data, nil
}

func (s *GaugeStorage) Increase(_ string, _ interface{}) (interface{}, error) {
	return nil, ErrNoMethod
}

func (s *GaugeStorage) GetLast(n string) (interface{}, error) {
	if _, ok := s.Storage[n]; !ok {
		return nil, ErrEmptyMetric
	}
	l := len(s.Storage[n])
	if l == 0 {
		return nil, ErrEmptyMetric
	}
	return s.Storage[n][l-1].Value, nil
}

func (s *GaugeStorage) GetLastAsString(n string) (string, error) {
	v, err := s.GetLast(n)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v), nil
}

func (s *GaugeStorage) GetAllLastAsString() (*MetricList, error) {
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
