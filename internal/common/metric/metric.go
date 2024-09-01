package metric

import (
	"screamer/internal/common"
)

type Type string

func (t *Type) String() string {
	return string(*t)
}

type Ident struct {
	Type Type
	Name string
}

type JsonMetric struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

type Metric struct {
	Ident Ident
	Value float64
}

func (m *Metric) Json() (JsonMetric, error) {
	switch m.Ident.Type {
	case Counter:
		return JsonMetric{
			ID:    m.Ident.Name,
			MType: m.Ident.Type.String(),
			Delta: nil,
			Value: &m.Value,
		}, nil
	case Gauge:
		v := int64(m.Value)
		return JsonMetric{
			ID:    m.Ident.Name,
			MType: m.Ident.Type.String(),
			Delta: &v,
			Value: nil,
		}, nil
	default:
		return JsonMetric{}, common.ErrTypeNotExists
	}
}

const (
	Counter Type = "counter"
	Gauge   Type = "gauge"
)

func newCounterIdent(n string) Ident {
	return Ident{
		Type: Counter,
		Name: n,
	}
}

func newGaugeIdent(n string) Ident {
	return Ident{
		Type: Gauge,
		Name: n,
	}
}

func NewMetricIdent(n string, t string) (Ident, error) {
	switch Type(t) {
	case Counter:
		return newCounterIdent(n), nil
	case Gauge:
		return newGaugeIdent(n), nil
	default:
		return Ident{}, common.ErrMetricNotExists
	}
}

func NewMetricIdentFromJson(jm *JsonMetric) (Ident, error) {
	switch Type(jm.MType) {
	case Counter:
		return newCounterIdent(jm.ID), nil
	case Gauge:
		return newGaugeIdent(jm.ID), nil
	default:
		return Ident{}, common.ErrMetricNotExists
	}
}

func NewCounter(n string, v float64) *Metric {
	return &Metric{
		Ident: newCounterIdent(n),
		Value: v,
	}
}

func NewGauge(n string, v float64) *Metric {
	return &Metric{
		Ident: newGaugeIdent(n),
		Value: v,
	}
}

func NewMetric(n string, v float64, t string) (*Metric, error) {
	switch Type(t) {
	case Counter:
		return NewCounter(n, v), nil
	case Gauge:
		return NewGauge(n, v), nil
	default:
		return nil, common.ErrMetricNotExists
	}
}

func NewMetricFromJson(jm *JsonMetric) (*Metric, error) {
	switch Type(jm.MType) {
	case Counter:
		return NewCounter(jm.ID, float64(*jm.Value)), nil
	case Gauge:
		return NewGauge(jm.ID, *jm.Value), nil
	default:
		return nil, common.ErrMetricNotExists
	}
}
