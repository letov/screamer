package metric

import (
	"screamer/internal/metric/validators"
)

type Ident int

const (
	CounterIdent Ident = iota
	GaugeIdent
)

var metricValidators = map[string]MetricValidator{
	validators.CounterLabel: {
		Ident:     CounterIdent,
		Validator: validators.CounterValidator,
	},
	validators.GaugeLabel: {
		Ident:     GaugeIdent,
		Validator: validators.GaugeValidator,
	},
}

type Raw struct {
	Label string
	Name  string
	Value string
}
type Metric struct {
	Ident Ident
	Name  string
	Value interface{}
}

type MetricValidator struct {
	Ident     Ident
	Validator func(value string) (interface{}, error)
}

func NewMetric(mr Raw) (Metric, error) {
	v, ok := metricValidators[mr.Label]
	if !ok {
		return Metric{}, validators.ErrUnknownMetricType
	}

	vv, err := v.Validator(mr.Value)
	if err != nil {
		return Metric{}, validators.ErrIncorrectMetricValue
	}

	return Metric{Ident: v.Ident, Name: mr.Name, Value: vv}, nil
}
