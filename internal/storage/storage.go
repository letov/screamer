package storage

import (
	"screamer/internal/metric"
	"screamer/internal/storage/repos"
)

type Storage interface {
	Add(m metric.Metric) error
	Get(k metric.Kind, n string) (interface{}, error)
	GetAsString(k metric.Kind, n string) (string, error)
	Debug() string
}

var storage Storage

func Init() {
	storage = repos.NewMemStorage()
}

func GetStorage() Storage {
	return storage
}
