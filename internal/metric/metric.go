package metric

import (
	"screamer/internal/metric/kinds"
)

type Kind int

const (
	Counter Kind = iota
	Gauge
)

var Validators = map[kinds.Label]Validator{
	kinds.CounterLabel: {
		Kind: Counter,
		Func: kinds.CounterValidator,
	},
	kinds.GaugeLabel: {
		Kind: Gauge,
		Func: kinds.GaugeValidator,
	},
}

type Raw struct {
	Label string
	Name  string
	Value string
}
type Metric struct {
	Kind  Kind
	Name  string
	Value interface{}
}

type Validator struct {
	Kind Kind
	Func func(value string) (interface{}, error)
}

func NewMetric(mr Raw) (Metric, error) {
	c, ok := Validators[kinds.Label(mr.Label)]
	if !ok {
		return Metric{}, kinds.ErrUnknownMetricType
	}

	v, err := c.Func(mr.Value)
	if err != nil {
		return Metric{}, kinds.ErrIncorrectMetricValue
	}

	return Metric{Kind: c.Kind, Name: mr.Name, Value: v}, nil
}

func LabelToKind(l string) (Kind, error) {
	switch kinds.Label(l) {
	case kinds.CounterLabel:
		return Counter, nil
	case kinds.GaugeLabel:
		return Gauge, nil
	default:
		return Counter, kinds.ErrUnknownMetricType
	}
}
