package config

import "screamer/internal/common/net-address"

type Config struct {
	NetAddress     net_address.NetAddress
	PollInterval   int
	ReportInterval int
	AgentLogEnable bool
}

type preConfig struct {
	NetAddress     *net_address.NetAddress
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
