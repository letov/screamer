package metric

import (
	"encoding/json"
	"screamer/internal/common"
	"strconv"
)

type Type string

func (t *Type) String() string {
	return string(*t)
}

type Ident struct {
	Type Type
	Name string
}

type JSONMetric struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

// Metric метрика, состоит из идентификатора и значения
type Metric struct {
	Ident Ident
	Value float64
}

func (m *Metric) JSON() (JSONMetric, error) {
	switch m.Ident.Type {
	case Counter:
		v := int64(m.Value)
		return JSONMetric{
			ID:    m.Ident.Name,
			MType: m.Ident.Type.String(),
			Delta: &v,
			Value: nil,
		}, nil
	case Gauge:
		return JSONMetric{
			ID:    m.Ident.Name,
			MType: m.Ident.Type.String(),
			Delta: nil,
			Value: &m.Value,
		}, nil
	default:
		return JSONMetric{}, common.ErrTypeNotExists
	}
}

func (m *Metric) Bytes() ([]byte, error) {
	jm, err := m.JSON()
	if err != nil {
		return []byte{}, err
	}

	return json.Marshal(jm)
}

func (m *Metric) String() string {
	switch m.Ident.Type {
	case Counter:
		return strconv.FormatFloat(m.Value, 'f', -1, 64)
	case Gauge:
		return strconv.FormatFloat(m.Value, 'g', 10, 64)
	default:
		return ""
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

func NewMetricIdentFromJSON(jm *JSONMetric) (Ident, error) {
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

func NewMetricFromJSON(jm *JSONMetric) (*Metric, error) {
	switch Type(jm.MType) {
	case Counter:
		if jm.Delta == nil {
			return nil, common.ErrEmptyValue
		}
		return NewCounter(jm.ID, float64(*jm.Delta)), nil
	case Gauge:
		if jm.Value == nil {
			return nil, common.ErrEmptyValue
		}
		return NewGauge(jm.ID, *jm.Value), nil
	default:
		return nil, common.ErrMetricNotExists
	}
}
