package services

import (
	"log"
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
		if err != nil && ps.config.AgentLogEnable {
			log.Println("Update metric error", err.Error())
		}
	}
}

func NewProcessingService(config *config.Config, repo repositories.Repository, metricSources []metric_sources.MetricSource) *ProcessingService {
	return &ProcessingService{
		config:        config,
		repo:          repo,
		metricSources: metricSources,
	}
}
