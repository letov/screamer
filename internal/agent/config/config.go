package config

import (
	"screamer/internal/agent/config/args"
	"screamer/internal/agent/config/base"
	"screamer/internal/agent/config/dotenv"
	"screamer/internal/agent/config/env"
	net_address "screamer/internal/common/net-address"
)

var config *base.Config

type EnvSrc struct {
	Args   *base.Config
	Env    *base.Config
	Dotenv *base.Config
}

func Init() {
	args.Init()
	env.Init()
	dotenv.Init()

	envSrc := EnvSrc{
		Args:   args.GetArgs(),
		Env:    env.GetEnv(),
		Dotenv: dotenv.GetDotenv(),
	}

	config = &base.Config{
		NetAddress:     getNetAddress(envSrc),
		PollInterval:   getPollInterval(envSrc),
		ReportInterval: getReportInterval(envSrc),
		AgentLogEnable: getAgentLogEnable(envSrc),
	}
}

func GetConfig() *base.Config {
	return config
}

func getNetAddress(envSrc EnvSrc) *net_address.NetAddress {
	if envSrc.Env.NetAddress != nil {
		return envSrc.Env.NetAddress
	} else if envSrc.Args.NetAddress != nil {
		return envSrc.Args.NetAddress
	}
	return envSrc.Dotenv.NetAddress
}

func getPollInterval(envSrc EnvSrc) *int {
	if envSrc.Env.PollInterval != nil {
		return envSrc.Env.PollInterval
	} else if envSrc.Args.PollInterval != nil {
		return envSrc.Args.PollInterval
	}
	return envSrc.Dotenv.PollInterval
}

func getReportInterval(envSrc EnvSrc) *int {
	if envSrc.Env.ReportInterval != nil {
		return envSrc.Env.ReportInterval
	} else if envSrc.Args.ReportInterval != nil {
		return envSrc.Args.ReportInterval
	}
	return envSrc.Dotenv.ReportInterval
}

func getAgentLogEnable(envSrc EnvSrc) *bool {
	if envSrc.Env.AgentLogEnable != nil {
		return envSrc.Env.AgentLogEnable
	} else if envSrc.Args.AgentLogEnable != nil {
		return envSrc.Args.AgentLogEnable
	}
	return envSrc.Dotenv.AgentLogEnable
}
