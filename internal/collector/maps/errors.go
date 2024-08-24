package maps

import "errors"

var TestSS = errors.New("cant typecast")
var ErrFloatTypecast = errors.New("cant typecast to float")
var ErrNotExists = errors.New("metric name not exists")
var ErrKindExists = errors.New("unknown metric kind")
