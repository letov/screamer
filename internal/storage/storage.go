package storage

import (
	"screamer/internal/metric"
	"screamer/internal/storage/repos"
)

type Storage interface {
	Add(m metric.Metric) error
	GetLast(k metric.Kind, n string) (interface{}, error)
	GetLastAsString(k metric.Kind, n string) (string, error)
}

var storage Storage

func Init() {
	storage = repos.NewMemStorage()
}

func GetStorage() Storage {
	return storage
}
