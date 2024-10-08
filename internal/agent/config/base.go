package config

import "screamer/internal/common/netaddress"

type Config struct {
	NetAddress     netaddress.NetAddress
	PollInterval   int
	ReportInterval int
	Key            string
	RateLimit      int
}

type preConfig struct {
	NetAddress     *netaddress.NetAddress
	PollInterval   *int
	ReportInterval *int
	Key            *string
	RateLimit      *int
}

type setConfig struct {
	NetAddress     bool
	PollInterval   bool
	ReportInterval bool
	Key            bool
	RateLimit      bool
}

func newPreConfig() preConfig {
	return preConfig{
		NetAddress:     nil,
		PollInterval:   nil,
		ReportInterval: nil,
		Key:            nil,
		RateLimit:      nil,
	}
}

func newSetConfig() setConfig {
	return setConfig{
		NetAddress:     false,
		PollInterval:   false,
		ReportInterval: false,
		Key:            false,
		RateLimit:      false,
	}
}
