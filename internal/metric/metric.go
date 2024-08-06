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
)

type Metric struct {
	Type  MetricType
	Name  string
	Value interface{}
}

func GetMetric(metricLabel, name, value string) (Metric, error) {
	switch metricLabel {
	case COUNTER_LABEL:
		v, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return Metric{}, err
		}
		return Metric{Type: COUNTER, Name: name, Value: v}, nil
	case GAUGE_LABEL:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return Metric{}, err
		}
		return Metric{Type: GAUGE, Name: name, Value: v}, nil
	}
	return Metric{}, errors.New("unknown metric type")
}
