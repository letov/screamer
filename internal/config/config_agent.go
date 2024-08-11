package config

import (
	"fmt"
	"screamer/internal/args"
)

var configAgent *ConfigAgent

type ConfigAgent struct {
	PollInterval   int
	ReportInterval int
	ServerURL      string
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

	var serverURL string
	if a.NetAddress.Host == "" {
		serverURL = getEnv("SERVER_URL", "http://localhost:8080")
	} else {
		serverURL = fmt.Sprintf("http://%v", a.NetAddress.String())
	}

	return &ConfigAgent{
		PollInterval:   pollInterval,
		ReportInterval: reportInterval,
		ServerURL:      serverURL,
		AgentLogEnable: getEnvInt("AGENT_LOG_ENABLE", 1) == 1,
	}
}
