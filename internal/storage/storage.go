package storage

type MemStorage struct {
	storage map[string]float64
}

func (s MemStorage) Add(name string, v float64) {
	s.storage[name] = v
}
