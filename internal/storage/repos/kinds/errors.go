package kinds

import "errors"

var ErrInvalidDataType = errors.New("invalid data type")
var ErrUnknownMetricaIdent = errors.New("unknown metrica ident")
var ErrEmptyMetric = errors.New("no values in metric")
var ErrNoMethod = errors.New("method no allowed")
