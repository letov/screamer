package maps

import "fmt"

type Gauge = map[string]float64

type GaugeMap struct {
	Map Gauge
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
