package services

import (
	"context"
	"encoding/json"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
	"sync"
)

type BackupService struct {
	config *config.Config
	repo   repositories.Repository
	log    *zap.SugaredLogger
	sync.Mutex
}

type JSONMetricList struct {
	Array []metric.JSONMetric
}

func (ps *BackupService) Save(ctx context.Context) {
	err := ps.toFile(ps.repo.GetAll(ctx))
	ps.processError(err)

	ps.log.Info("Saved backup")
}

func (ps *BackupService) Load(ctx context.Context) {
	ms, err := ps.fromFile()
	if err != nil {
		ps.processError(err)
		return
	}

	for _, m := range ms {
		_, err = ps.repo.Add(ctx, *m)
		ps.processError(err)
	}

	ps.log.Info("Loaded backup")
}

func (ps *BackupService) toFile(ms []metric.Metric) error {
	fp := ps.config.FileStoragePath

	jms := make([]metric.JSONMetric, 0)
	for _, m := range ms {
		j, err := m.JSON()
		ps.processError(err)
		jms = append(jms, j)
	}

	jml := &JSONMetricList{Array: jms}
	body, err := json.MarshalIndent(jml, "", "   ")
	ps.processError(err)

	ps.Lock()
	err = os.WriteFile(fp, body, 0777)
	ps.Unlock()
	return err
}

func (ps *BackupService) fromFile() ([]*metric.Metric, error) {
	fp := ps.config.FileStoragePath

	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	jml := &JSONMetricList{}
	err = json.Unmarshal(data, jml)
	if err != nil {
		return nil, err
	}

	res := make([]*metric.Metric, 0)
	for _, jm := range jml.Array {
		m, err := metric.NewMetricFromJSON(&jm)
		ps.processError(err)
		res = append(res, m)
	}

	return res, nil
}

func (ps *BackupService) processError(err error) {
	if err != nil {
		ps.log.Warn("Save backup error:", err.Error())
	}
}

func NewBackupService(lc fx.Lifecycle, log *zap.SugaredLogger, c *config.Config, r repositories.Repository) *BackupService {
	srv := &BackupService{
		config: c,
		repo:   r,
		log:    log,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Loading backup")
			srv.Load(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Saving backup")
			srv.Save(ctx)
			return nil
		},
	})

	return srv
}
