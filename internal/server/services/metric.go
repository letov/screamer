package services

import (
	"context"
	"encoding/json"
	"fmt"
	"screamer/internal/common"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
	"strconv"
	"sync"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MetricService struct {
	config *config.Config
	repo   repositories.Repository
	stop   bool
	sync.Mutex
	activeWrites int
}

func (ms *MetricService) increaseActiveWrites() {
	ms.Lock()
	ms.activeWrites++
	ms.Unlock()
}

func (ms *MetricService) decreaseActiveWrites() {
	ms.Lock()
	ms.activeWrites--
	ms.Unlock()
}

func (ms *MetricService) UpdatesMetricJSON(ctx context.Context, body *[]byte) (err error) {
	if ms.stop {
		return common.ErrServiceStop
	}
	ms.increaseActiveWrites()
	defer ms.decreaseActiveWrites()

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

func (ms *MetricService) UpdateMetricJSON(ctx context.Context, body *[]byte) (res *[]byte, err error) {
	var jm metric.JSONMetric
	err = json.Unmarshal(*body, &jm)
	if err != nil {
		return nil, err
	}

	m, err := metric.NewMetricFromJSON(&jm)
	if err != nil {
		return nil, err
	}

	return ms.processUpdateMetric(ctx, m)
}

func (ms *MetricService) UpdateMetricParams(ctx context.Context, n string, vs string, t string) (res *[]byte, err error) {
	v, err := strconv.ParseFloat(vs, 64)
	if err != nil {
		return nil, err
	}

	m, err := metric.NewMetric(n, v, t)
	if err != nil {
		return nil, err
	}

	return ms.processUpdateMetric(ctx, m)
}

func (ms *MetricService) processUpdateMetric(ctx context.Context, m *metric.Metric) (res *[]byte, err error) {
	if ms.stop {
		return nil, common.ErrServiceStop
	}
	ms.increaseActiveWrites()
	defer ms.decreaseActiveWrites()

	var newM metric.Metric

	if m.Ident.Type == metric.Counter {
		newM, err = ms.repo.Increase(ctx, *m)
	} else {
		newM, err = ms.repo.Add(ctx, *m)
	}
	if err != nil {
		return nil, err
	}

	body, err := newM.Bytes()
	return &body, err
}

func (ms *MetricService) ValueMetricJSON(ctx context.Context, body *[]byte) (res *[]byte, err error) {
	var jm metric.JSONMetric
	err = json.Unmarshal(*body, &jm)
	if err != nil {
		return nil, err
	}

	i, err := metric.NewMetricIdentFromJSON(&jm)
	if err != nil {
		return nil, err
	}

	m, err := ms.repo.Get(ctx, i)
	if err != nil {
		return nil, err
	}

	bs, err := m.Bytes()
	if err != nil {
		return nil, err
	}

	return &bs, err
}

func (ms *MetricService) ValueMetricParams(ctx context.Context, n string, t string) (res *[]byte, err error) {
	i, err := metric.NewMetricIdent(n, t)
	if err != nil {
		return nil, err
	}

	m, err := ms.repo.Get(ctx, i)
	bs := []byte(m.String())

	return &bs, err
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
		config:       c,
		repo:         r,
		stop:         false,
		activeWrites: 0,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			ms.stop = true
			log.Info("Service active writes count: ", ms.activeWrites)
			if ms.activeWrites == 0 {
				log.Info("Service closed")
				return nil
			}
			for i := 0; i < 5; i++ {
				if ms.activeWrites > 0 {
					log.Info("Service active writes count: ", ms.activeWrites)
					log.Info("Try to wait writing")
					time.Sleep(time.Second)
				} else {
					log.Info("Try to wait")
					break
				}
			}
			return nil
		},
	})

	return ms
}
