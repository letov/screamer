package services

import (
	"context"
	"screamer/internal/agent/config"
	metric_sources "screamer/internal/agent/metricsources"
	"screamer/internal/agent/repositories"
	"screamer/internal/common"
	"screamer/internal/common/metric"

	"go.uber.org/zap"
)

type ProcessingService struct {
	config        *config.Config
	repo          repositories.Repository
	metricSources []metric_sources.MetricSource
	log           *zap.SugaredLogger
}

func (ps *ProcessingService) UpdateMetrics(ctx context.Context) {
	ms := make([]*metric.Metric, 0)
	for _, fn := range ps.metricSources {
		ms = append(ms, fn()...)
	}
	for _, m := range ms {
		var err error
		switch m.Ident.Type {
		case metric.Counter:
			_, err = ps.repo.Increase(ctx, m.Ident, 1)
		case metric.Gauge:
			_, err = ps.repo.Update(ctx, *m)
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
