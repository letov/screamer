package config

import "screamer/internal/args"

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

	var serverUrl string
	if a.NetAddress.Port == UnsetIntValue {
		serverUrl = getEnv("SERVER_URL", "http://localhost:8080")
	} else {
		serverUrl = a.NetAddress.String()
	}

	var pollInterval int
	if a.NetAddress.Port == UnsetIntValue {
		pollInterval = getEnvInt("POLL_INTERVAL", 2)
	} else {
		pollInterval = *a.PollInterval
	}

	var reportInterval int
	if a.NetAddress.Port == UnsetIntValue {
		reportInterval = getEnvInt("REPORT_INTERVAL", 10)
	} else {
		reportInterval = *a.ReportInterval
	}

	return &ConfigAgent{
		PollInterval:   pollInterval,
		ReportInterval: reportInterval,
		ServerUrl:      serverUrl,
		AgentLogEnable: getEnvInt("AGENT_LOG_ENABLE", 1) == 1,
	}
}
