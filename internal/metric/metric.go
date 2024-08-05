package metric

import (
	"errors"
	"strconv"
)

const (
	COUNTER_LABEL string = "counter"
	GAUGE_LABEL          = "gauge"
)

type MetricType int

const (
	COUNTER MetricType = iota
	GAUGE
	EMPTY
)

type Metric struct {
	Type  MetricType
	Name  string
	Value float64
}

func GetMetric(metricLabel, name, value string) (Metric, error) {
	switch metricLabel {
	case COUNTER_LABEL:
		v, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return Metric{Type: EMPTY, Name: name, Value: float64(v)}, err
		}
		return Metric{Type: COUNTER, Name: name, Value: float64(0)}, nil
	case GAUGE_LABEL:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return Metric{Type: EMPTY, Name: name, Value: v}, err
		}
		return Metric{Type: GAUGE, Name: name, Value: float64(0)}, nil
	}
	return Metric{Type: EMPTY, Name: name, Value: float64(0)}, errors.New("unknown metric type")
}
