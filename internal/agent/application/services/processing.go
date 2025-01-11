package services

import (
	"context"
	"screamer/internal/agent/infrastructure/config"
	metric_sources "screamer/internal/agent/infrastructure/metricsources"
	"screamer/internal/agent/infrastructure/repositories"
	"screamer/internal/common"
	"screamer/internal/common/domain/metric"

	"go.uber.org/zap"
)

type Processing struct {
	config        *config.Config
	repo          repositories.Repository
	metricSources []metric_sources.MetricSource
	log           *zap.SugaredLogger
}

func (ps *Processing) UpdateMetrics(ctx context.Context) {
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

func NewProcessing(log *zap.SugaredLogger, config *config.Config, repo repositories.Repository, metricSources []metric_sources.MetricSource) *Processing {
	return &Processing{
		config:        config,
		repo:          repo,
		metricSources: metricSources,
		log:           log,
	}
}
