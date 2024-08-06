package metric

import (
	"errors"
	"strconv"
)

type Label string
type Kind int

const (
	CounterLabel Label = "counter"
	GougeLabel   Label = "gauge"
)
const (
	Counter Kind = iota
	Gauge
)

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

var ErrUnknownMetricType = errors.New("unknown metric type")
var ErrIncorrectMetricValue = errors.New("incorrect metric value")

func NewMetric(mr Raw) (Metric, error) {
	switch Label(mr.Label) {
	case CounterLabel:
		v, err := strconv.ParseInt(mr.Value, 0, 64)
		if err != nil {
			return Metric{}, ErrIncorrectMetricValue
		}
		return Metric{Kind: Counter, Name: mr.Name, Value: v}, nil
	case GougeLabel:
		v, err := strconv.ParseFloat(mr.Value, 64)
		if err != nil {
			return Metric{}, ErrIncorrectMetricValue
		}
		return Metric{Kind: Gauge, Name: mr.Name, Value: v}, nil
	}
	return Metric{}, ErrUnknownMetricType
}
