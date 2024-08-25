package storage

import (
	"screamer/internal/metric"
	"screamer/internal/storage/repos"
	"screamer/internal/storage/repos/kinds"
)

type Storage interface {
	Add(m metric.Metric) (interface{}, error)
	GetLast(k metric.Kind, n string) (interface{}, error)
	GetLastAsString(k metric.Kind, n string) (string, error)
	GetAllLastAsString(k metric.Kind) (*kinds.MetricList, error)
}

var storage Storage

func Init() {
	storage = repos.NewMemStorage()
}

func GetStorage() Storage {
	return storage
}
