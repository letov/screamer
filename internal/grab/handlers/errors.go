package handlers

import "errors"

var ErrNoStorage = errors.New("there is no init storage")
var ErrNoMetric = errors.New("metric not exists")
