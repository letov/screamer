package config

import (
	"screamer/internal/common/helpers/netaddress"
)

type Config struct {
	NetAddress     netaddress.NetAddress
	PollInterval   int
	ReportInterval int
	Key            string
	RateLimit      int
	CryptoKey      string
	Host           string
	NetAddressGrpc netaddress.NetAddress
}

type preConfig struct {
	NetAddress     *netaddress.NetAddress
	PollInterval   *int
	ReportInterval *int
	Key            *string
	RateLimit      *int
	CryptoKey      *string
	Host           *string
	NetAddressGrpc *netaddress.NetAddress
}

type setConfig struct {
	NetAddress     bool
	PollInterval   bool
	ReportInterval bool
	Key            bool
	RateLimit      bool
	CryptoKey      bool
	Host           bool
	NetAddressGrpc bool
}

func newPreConfig() preConfig {
	return preConfig{
		NetAddress:     nil,
		PollInterval:   nil,
		ReportInterval: nil,
		Key:            nil,
		RateLimit:      nil,
		CryptoKey:      nil,
		Host:           nil,
		NetAddressGrpc: nil,
	}
}

func newSetConfig() setConfig {
	return setConfig{
		NetAddress:     false,
		PollInterval:   false,
		ReportInterval: false,
		Key:            false,
		RateLimit:      false,
		CryptoKey:      false,
		Host:           false,
		NetAddressGrpc: false,
	}
}
