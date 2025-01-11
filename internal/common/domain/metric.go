package domain

import (
	"screamer/internal/common"
	"strconv"
)

// Type тип метрики
type Type string

func (t Type) String() string {
	return string(t)
}

const (
	Counter Type = "counter"
	Gauge   Type = "gauge"
)

// Ident идентификатор метрики из типа и имени
type Ident struct {
	Type Type
	Name string
}

// Metric метрика - идентификатора и значения
type Metric struct {
	Ident Ident
	Value float64
}

func (m Metric) String() string {
	switch m.Ident.Type {
	case Counter:
		return strconv.FormatFloat(m.Value, 'f', -1, 64)
	case Gauge:
		return strconv.FormatFloat(m.Value, 'g', 10, 64)
	default:
		return ""
	}
}

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

func NewMetricIdent(n string, t Type) (Ident, error) {
	switch t {
	case Counter:
		return newCounterIdent(n), nil
	case Gauge:
		return newGaugeIdent(n), nil
	default:
		return Ident{}, common.ErrMetricNotExists
	}
}

func newCounter(n string, v float64) Metric {
	return Metric{
		Ident: newCounterIdent(n),
		Value: v,
	}
}

func newGauge(n string, v float64) Metric {
	return Metric{
		Ident: newGaugeIdent(n),
		Value: v,
	}
}

func NewMetric(n string, v float64, t Type) (Metric, error) {
	switch t {
	case Counter:
		return newCounter(n, v), nil
	case Gauge:
		return newGauge(n, v), nil
	default:
		return Metric{}, common.ErrMetricNotExists
	}
}
