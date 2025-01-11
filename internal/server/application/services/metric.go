package services

import (
	"context"
	"fmt"
	"screamer/internal/common/application/dto"
	"screamer/internal/common/domain"
	"screamer/internal/server/application/repo"
	"screamer/internal/server/infrastructure/config"
	"sync/atomic"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MetricService struct {
	config     *config.Config
	repo       repo.Repository
	stop       bool
	activeJobs atomic.Int32
}

func (ms *MetricService) UpdateBatchMetricJSON(ctx context.Context, jms []dto.JSONMetric) (err error) {
	ms.activeJobs.Add(1)
	defer func() {
		ms.activeJobs.Add(-1)
	}()

	mList := make([]domain.Metric, 0)

	for _, jm := range jms {
		m, err := jm.GetDomainMetric()
		if err != nil {
			return err
		}
		mList = append(mList, m)
	}

	return ms.repo.BatchUpdate(ctx, mList)
}

func (ms *MetricService) UpdateMetricJSON(ctx context.Context, jm dto.JSONMetric) (dto.JSONMetric, error) {
	m, err := jm.GetDomainMetric()
	if err != nil {
		return dto.JSONMetric{}, err
	}

	return ms.processUpdateMetric(ctx, m)
}

func (ms *MetricService) processUpdateMetric(ctx context.Context, m domain.Metric) (dto.JSONMetric, error) {
	ms.activeJobs.Add(1)
	defer func() {
		ms.activeJobs.Add(-1)
	}()

	var res domain.Metric
	var e error
	if m.Ident.Type == domain.Counter {
		res, e = ms.repo.Increase(ctx, m)
	} else {
		res, e = ms.repo.Add(ctx, m)
	}
	if e != nil {
		return dto.JSONMetric{}, e
	}
	return dto.NewJSONMetric(res)
}

func (ms *MetricService) ValueMetricJSON(ctx context.Context, jm dto.JSONMetric) (domain.Metric, error) {
	i, err := jm.GetIdent()
	if err != nil {
		return domain.Metric{}, err
	}

	return ms.repo.Get(ctx, i)
}

func (ms *MetricService) Home(ctx context.Context) (res *[]byte) {
	r := "<html><body>"
	r += "<h1>Metrics</h1>"
	metrics := ms.repo.GetAll(ctx)
	for _, m := range metrics {
		r += fmt.Sprintf("<p>%v: %v</p>", m.Ident.Name, m.Value)
	}
	r += "</body></html>"

	bs := []byte(r)

	return &bs
}

func NewMetricService(
	lc fx.Lifecycle,
	log *zap.SugaredLogger,
	c *config.Config,
	r repo.Repository,
) *MetricService {
	ms := &MetricService{
		config: c,
		repo:   r,
		stop:   false,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			aj := ms.activeJobs.Load()
			log.Info("Server active jobs count: ", aj)
			if aj != 0 {
				for i := 0; i < 5; i++ {
					aj = ms.activeJobs.Load()
					if aj > 0 {
						log.Info("Server active jobs count: ", aj)
						log.Info("Try to wait")
						time.Sleep(time.Second)
					} else {
						break
					}
				}
			}
			log.Info("Sending service closed")
			return nil
		},
	})

	return ms
}
