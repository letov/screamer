package config

import (
	"fmt"
	"screamer/internal/args"
)

var configAgent *ConfigAgent

type ConfigAgent struct {
	PollInterval   int
	ReportInterval int
	ServerUrl      string
	AgentLogEnable bool
}

func InitAgent() {
	Init()
	configAgent = newConfigAgent()
}

func GetConfigAgent() *ConfigAgent {
	return configAgent
}

func newConfigAgent() *ConfigAgent {
	a := args.GetArgsAgent()

	var pollInterval int
	if *a.PollInterval == 0 {
		pollInterval = getEnvInt("POLL_INTERVAL", 2)
	} else {
		pollInterval = *a.PollInterval
	}

	var reportInterval int
	if *a.ReportInterval == 0 {
		reportInterval = getEnvInt("REPORT_INTERVAL", 10)
	} else {
		reportInterval = *a.ReportInterval
	}

	var serverUrl string
	if a.NetAddress.Host == "" {
		serverUrl = getEnv("SERVER_URL", "http://localhost:8080")
	} else {
		serverUrl = fmt.Sprintf("http://%v", a.NetAddress.String())
	}

	return &ConfigAgent{
		PollInterval:   pollInterval,
		ReportInterval: reportInterval,
		ServerUrl:      serverUrl,
		AgentLogEnable: getEnvInt("AGENT_LOG_ENABLE", 1) == 1,
	}
}
