package services

import (
	"encoding/json"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
	"strconv"
)

type MetricService struct {
	config        *config.Config
	repo          repositories.Repository
	backupService *BackupService
}

func (ms *MetricService) UpdateMetricJSON(body *[]byte) (res *[]byte, err error) {
	var jm metric.JSONMetric
	err = json.Unmarshal(*body, &jm)
	if err != nil {
		return nil, err
	}

	m, err := metric.NewMetricFromJSON(&jm)
	if err != nil {
		return nil, err
	}

	return ms.processMetric(m)
}

func (ms *MetricService) UpdateMetricParams(n string, vs string, t string) (res *[]byte, err error) {
	v, err := strconv.ParseFloat(vs, 64)
	if err != nil {
		return nil, err
	}

	m, err := metric.NewMetric(n, v, t)
	if err != nil {
		return nil, err
	}

	return ms.processMetric(m)
}

func (ms *MetricService) processMetric(m *metric.Metric) (res *[]byte, err error) {
	var newM metric.Metric

	if m.Ident.Type == metric.Counter {
		newM, err = ms.repo.Increase(*m)
	} else {
		newM, err = ms.repo.Add(*m)
	}
	if err != nil {
		return nil, err
	}

	if ms.config.Restore && ms.config.StoreInterval == 0 {
		ms.backupService.Save()
	}

	body, err := newM.Bytes()
	return &body, err
}

func NewMetricService(c *config.Config, r repositories.Repository, bs *BackupService) *MetricService {
	return &MetricService{
		config:        c,
		repo:          r,
		backupService: bs,
	}
}
