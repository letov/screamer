package common

import "errors"

var ErrInvalidAddr = errors.New("need address in a form host:port")
var ErrMetricNotExists = errors.New("metric not exists")
var ErrTypeNotExists = errors.New("unknown metric type")
var ErrInvalidValue = errors.New("invalid value")
var ErrEmptyValue = errors.New("empty value")
