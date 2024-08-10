package collector_maps

import "fmt"

type Counter = map[string]int64

type CounterMap struct {
	Map Counter
}

func NewCounterMap(init *[]string) *CounterMap {
	initMap := make(Counter)

	for _, l := range *init {
		initMap[l] = 0
	}

	return &CounterMap{
		Map: initMap,
	}
}

func (m *CounterMap) Update(n string, v interface{}) error {
	f, ok := v.(int64)
	if !ok {
		return ErrFloatTypecast
	}
	m.Map[n] = f
	return nil
}

func (m *CounterMap) Get(n string) (interface{}, error) {
	v, ok := m.Map[n]
	if !ok {
		return nil, ErrNotExists
	}
	return v, nil
}

func (m *CounterMap) Export() map[string]string {
	res := make(map[string]string)
	for n, v := range m.Map {
		res[n] = fmt.Sprintf("%d", v)
	}
	return res
}
