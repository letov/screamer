package validators

import (
	"strconv"
)

const GaugeLabel string = "gauge"

func GaugeValidator(value string) (interface{}, error) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, ErrIncorrectMetricValue
	}
	return v, nil
}
