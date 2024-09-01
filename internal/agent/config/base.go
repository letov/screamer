package config

import "screamer/internal/common/netaddress"

type Config struct {
	NetAddress     netaddress.NetAddress
	PollInterval   int
	ReportInterval int
	AgentLogEnable bool
}

type preConfig struct {
	NetAddress     *netaddress.NetAddress
	PollInterval   *int
	ReportInterval *int
	AgentLogEnable *bool
}

type setConfig struct {
	NetAddress     bool
	PollInterval   bool
	ReportInterval bool
	AgentLogEnable bool
}

func newPreConfig() preConfig {
	return preConfig{
		NetAddress:     nil,
		PollInterval:   nil,
		ReportInterval: nil,
		AgentLogEnable: nil,
	}
}

func newSetConfig() setConfig {
	return setConfig{
		NetAddress:     false,
		PollInterval:   false,
		ReportInterval: false,
		AgentLogEnable: false,
	}
}
