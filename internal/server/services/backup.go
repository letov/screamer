package services

import (
	"encoding/json"
	"log"
	"os"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
	"sync"
)

type BackupService struct {
	config *config.Config
	repo   repositories.Repository
	sync.Mutex
}

type JsonMetricList struct {
	Array []metric.JsonMetric
}

func (ps *BackupService) Save() {
	ps.Lock()
	err := ps.toFile(ps.repo.GetAll())
	ps.processError(err)
	ps.Unlock()
}

func (ps *BackupService) Load() {
	ms, err := ps.fromFile()
	if err != nil {
		ps.processError(err)
		return
	}

	ps.Lock()
	for _, m := range ms {
		_, err = ps.repo.Add(*m)
		ps.processError(err)
	}
	ps.Unlock()
}

func (ps *BackupService) toFile(ms []metric.Metric) error {
	fp := ps.config.FileStoragePath

	jms := make([]metric.JsonMetric, 0)
	for _, m := range ms {
		j, err := m.Json()
		ps.processError(err)
		jms = append(jms, j)
	}

	jml := &JsonMetricList{Array: jms}
	body, err := json.MarshalIndent(jml, "", "   ")
	ps.processError(err)

	return os.WriteFile(fp, body, 0777)
}

func (ps *BackupService) fromFile() ([]*metric.Metric, error) {
	fp := ps.config.FileStoragePath

	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	jml := &JsonMetricList{}
	err = json.Unmarshal(data, jml)
	if err != nil {
		return nil, err
	}

	res := make([]*metric.Metric, 0)
	for _, jm := range jml.Array {
		m, err := metric.NewMetricFromJson(&jm)
		ps.processError(err)
		res = append(res, m)
	}

	return res, nil
}

func (ps *BackupService) processError(err error) {
	if err != nil && ps.config.ServerLogEnable {
		log.Println("Save backup error:", err.Error())
	}
}

func NewBackupService(c *config.Config, r repositories.Repository) *BackupService {
	return &BackupService{
		config: c,
		repo:   r,
	}
}
