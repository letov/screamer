package storage

import (
	"screamer/internal/metric"
	"screamer/internal/storage/repos"
)

type Storage interface {
	Add(m metric.Metric) error
	Debug() string
}

var storage Storage

func Init() {
	storage = repos.NewMemStorage()
}

func GetStorage() Storage {
	return storage
}
