package agentservices

import (
	"context"
	"screamer/internal/agent/application/repo"
	"screamer/internal/agent/infrastructure/config"
	metric_sources "screamer/internal/agent/infrastructure/metricsources"
	"screamer/internal/common"
	"screamer/internal/common/domain"

	"go.uber.org/zap"
)

type Processing struct {
	config        *config.Config
	repo          repo.Repository
	metricSources []metric_sources.MetricSource
	log           *zap.SugaredLogger
}

func (ps *Processing) UpdateMetrics(ctx context.Context) {
	ms := make([]domain.Metric, 0)
	for _, fn := range ps.metricSources {
		ms = append(ms, fn()...)
	}
	for _, m := range ms {
		var err error
		switch m.Ident.Type {
		case domain.Counter:
			_, err = ps.repo.Increase(ctx, m.Ident, 1)
		case domain.Gauge:
			_, err = ps.repo.Update(ctx, m)
		default:
			err = common.ErrTypeNotExists
		}
		if err != nil {
			ps.log.Warn("Update metric error", err.Error())
		}
	}
}

func NewProcessing(log *zap.SugaredLogger, config *config.Config, repo repo.Repository, metricSources []metric_sources.MetricSource) *Processing {
	return &Processing{
		config:        config,
		repo:          repo,
		metricSources: metricSources,
		log:           log,
	}
}
