package config

import "screamer/internal/common/netaddress"

type Config struct {
	NetAddress     netaddress.NetAddress
	PollInterval   int
	ReportInterval int
	Key            string
}

type preConfig struct {
	NetAddress     *netaddress.NetAddress
	PollInterval   *int
	ReportInterval *int
	Key            *string
}

type setConfig struct {
	NetAddress     bool
	PollInterval   bool
	ReportInterval bool
	Key            bool
}

func newPreConfig() preConfig {
	return preConfig{
		NetAddress:     nil,
		PollInterval:   nil,
		ReportInterval: nil,
		Key:            nil,
	}
}

func newSetConfig() setConfig {
	return setConfig{
		NetAddress:     false,
		PollInterval:   false,
		ReportInterval: false,
		Key:            false,
	}
}
