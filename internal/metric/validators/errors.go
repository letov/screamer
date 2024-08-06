package validators

import "errors"

var ErrUnknownMetricType = errors.New("unknown metric type")
var ErrIncorrectMetricValue = errors.New("incorrect metric value")
