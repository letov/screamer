package maps

import (
	"fmt"
	"screamer/internal/metric/kinds"
)

type Gauge = map[string]float64

type GaugeMap struct {
	Map Gauge
}

func (m *GaugeMap) ExportJsonMetrics() []*JsonMetric {
	res := make([]*JsonMetric, 0)
	for n := range m.Map {
		el, err := m.GetJsonMetric(n)
		if err == nil {
			res = append(res, el)
		}
	}
	return res
}

func (m *GaugeMap) GetJsonMetric(n string) (*JsonMetric, error) {
	v, ok := m.Map[n]
	if !ok {
		return nil, ErrNotExists
	}
	return &JsonMetric{
		ID:    n,
		MType: string(kinds.GaugeLabel),
		Delta: nil,
		Value: &v,
	}, nil
}

func NewGaugeMap() *GaugeMap {
	return &GaugeMap{
		Map: make(Gauge),
	}
}

func (m *GaugeMap) Update(n string, v interface{}) error {
	f, ok := v.(float64)
	if !ok {
		return ErrFloatTypecast
	}
	m.Map[n] = f
	return nil
}

func (m *GaugeMap) Get(n string) (interface{}, error) {
	v, ok := m.Map[n]
	if !ok {
		return nil, ErrNotExists
	}
	return v, nil
}

func (m *GaugeMap) Export() map[string]string {
	res := make(map[string]string)
	for n, v := range m.Map {
		res[n] = fmt.Sprintf("%f", v)
	}
	return res
}
