package common

import "errors"

var (
	ErrInvalidAddr     = errors.New("need address in a form host:port")
	ErrMetricNotExists = errors.New("metric not exists")
	ErrTypeNotExists   = errors.New("unknown metric type")
	ErrInvalidValue    = errors.New("invalid value")
	ErrEmptyValue      = errors.New("empty value")

	ErrNoOKResponse   = errors.New("no ok response")
	ErrNoDBConnection = errors.New("no db connection")

	ErrInvalidHash = errors.New("invalid hash")

	ErrServiceStop = errors.New("service is stop")
)
