package services

import (
	"go.uber.org/zap"
	"screamer/internal/agent/config"
	metric_sources "screamer/internal/agent/metricsources"
	"screamer/internal/agent/repositories"
	"screamer/internal/common"
	"screamer/internal/common/metric"
)

type ProcessingService struct {
	config        *config.Config
	repo          repositories.Repository
	metricSources []metric_sources.MetricSource
	log           *zap.SugaredLogger
}

func (ps *ProcessingService) UpdateMetrics() {
	ms := make([]*metric.Metric, 0)
	for _, fn := range ps.metricSources {
		ms = append(ms, fn()...)
	}
	for _, m := range ms {
		var err error
		switch m.Ident.Type {
		case metric.Counter:
			_, err = ps.repo.Increase(m.Ident, 1)
		case metric.Gauge:
			_, err = ps.repo.Update(*m)
		default:
			err = common.ErrTypeNotExists
		}
		if err != nil {
			ps.log.Warn("Update metric error", err.Error())
		}
	}
}

func NewProcessingService(log *zap.SugaredLogger, config *config.Config, repo repositories.Repository, metricSources []metric_sources.MetricSource) *ProcessingService {
	return &ProcessingService{
		config:        config,
		repo:          repo,
		metricSources: metricSources,
		log:           log,
	}
}
