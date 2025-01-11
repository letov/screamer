package services

import (
	"context"
	"encoding/json"
	"fmt"
	"screamer/internal/common/domain/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
	"strconv"
	"sync/atomic"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MetricService struct {
	config     *config.Config
	repo       repositories.Repository
	stop       bool
	activeJobs atomic.Int32
}

func (ms *MetricService) UpdateBatchMetricJSON(ctx context.Context, body *[]byte) (err error) {
	ms.activeJobs.Add(1)
	defer func() {
		ms.activeJobs.Add(-1)
	}()

	var jms []metric.JSONMetric
	err = json.Unmarshal(*body, &jms)
	if err != nil {
		return
	}

	mList := make([]metric.Metric, 0)

	for _, jm := range jms {
		m, err := metric.NewMetricFromJSON(&jm)
		if err != nil {
			return err
		}
		mList = append(mList, *m)
	}

	return ms.repo.BatchUpdate(ctx, mList)
}

func (ms *MetricService) UpdateMetricJSON(ctx context.Context, jm metric.JSONMetric) (metric.Metric, error) {
	m, err := metric.NewMetricFromJSON(&jm)
	if err != nil {
		return metric.Metric{}, err
	}

	return ms.processUpdateMetric(ctx, m)
}

func (ms *MetricService) UpdateMetricParams(ctx context.Context, n string, vs string, t string) (metric.Metric, error) {
	v, err := strconv.ParseFloat(vs, 64)
	if err != nil {
		return metric.Metric{}, err
	}

	m, err := metric.NewMetric(n, v, t)
	if err != nil {
		return metric.Metric{}, err
	}

	return ms.processUpdateMetric(ctx, m)
}

func (ms *MetricService) processUpdateMetric(ctx context.Context, m *metric.Metric) (metric.Metric, error) {
	ms.activeJobs.Add(1)
	defer func() {
		ms.activeJobs.Add(-1)
	}()

	if m.Ident.Type == metric.Counter {
		return ms.repo.Increase(ctx, *m)
	} else {
		return ms.repo.Add(ctx, *m)
	}
}

func (ms *MetricService) ValueMetricJSON(ctx context.Context, jm metric.JSONMetric) (metric.Metric, error) {
	i, err := metric.NewMetricIdentFromJSON(&jm)
	if err != nil {
		return metric.Metric{}, err
	}

	return ms.repo.Get(ctx, i)
}

func (ms *MetricService) ValueMetricParams(ctx context.Context, n string, t string) (metric.Metric, error) {
	i, err := metric.NewMetricIdent(n, t)
	if err != nil {
		return metric.Metric{}, err
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
	r repositories.Repository,
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
