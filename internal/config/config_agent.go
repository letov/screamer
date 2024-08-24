package config

import (
	"fmt"
	"screamer/internal/config/args"
	"screamer/internal/config/dotenv"
	"screamer/internal/config/env"
)

var configA *ConfigA

type ConfigA struct {
	PollInterval   int
	ReportInterval int
	ServerURL      string
	AgentLogEnable bool
}

type EnvSrcA struct {
	ArgsA   *args.ArgsA
	EnvA    *env.EnvA
	DotenvA *dotenv.DotenvA
}

func InitAgent() {
	args.InitAgent()
	env.InitAgent()
	dotenv.InitAgent()

	envSrcA := EnvSrcA{
		ArgsA:   args.GetArgsA(),
		EnvA:    env.GetEnvA(),
		DotenvA: dotenv.GetDotenvA(),
	}

	configA = &ConfigA{
		PollInterval:   getPollInterval(&envSrcA),
		ReportInterval: getReportInterval(&envSrcA),
		ServerURL:      getServerURL(&envSrcA),
		AgentLogEnable: envSrcA.DotenvA.AgentLogEnable,
	}
}

func GetConfigA() *ConfigA {
	return configA
}

func getPollInterval(envSrcA *EnvSrcA) int {
	if envSrcA.EnvA.PollInterval != 0 {
		return envSrcA.EnvA.PollInterval
	} else if *envSrcA.ArgsA.PollInterval != 0 {
		return *envSrcA.ArgsA.PollInterval
	}
	return envSrcA.DotenvA.PollInterval
}

func getReportInterval(envSrcA *EnvSrcA) int {
	if envSrcA.EnvA.ReportInterval != 0 {
		return envSrcA.EnvA.ReportInterval
	} else if *envSrcA.ArgsA.ReportInterval != 0 {
		return *envSrcA.ArgsA.ReportInterval
	}
	return envSrcA.DotenvA.ReportInterval
}

func getServerURL(envSrcA *EnvSrcA) string {
	var serverURL string
	if envSrcA.EnvA.Address != "" {
		serverURL = envSrcA.EnvA.Address
	} else if envSrcA.ArgsA.NetAddress.Host != "" {
		serverURL = envSrcA.ArgsA.NetAddress.String()
	} else {
		serverURL = envSrcA.DotenvA.Address
	}
	return fmt.Sprintf("http://%v", serverURL)
}
