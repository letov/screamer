package kinds

import (
	"strconv"
)

const CounterLabel string = "counter"

func CounterValidator(value string) (interface{}, error) {
	v, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return 0, ErrIncorrectMetricValue
	}
	return v, nil
}
