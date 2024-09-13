package services

import (
	"context"
	"encoding/json"
	"fmt"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"screamer/internal/server/repositories"
	"strconv"
)

type MetricService struct {
	config *config.Config
	repo   repositories.Repository
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

func NewMetricService(c *config.Config, r repositories.Repository) *MetricService {
	return &MetricService{
		config: c,
		repo:   r,
	}
}
