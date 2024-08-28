package env

import (
	"os"
	"screamer/internal/agent/config/base"
	net_address "screamer/internal/common/net-address"
	"strconv"
)

var env *base.Config

func Init() {
	env = &base.Config{
		NetAddress:     nil,
		PollInterval:   nil,
		ReportInterval: nil,
		AgentLogEnable: nil,
	}

	a, exists := os.LookupEnv("ADDRESS")
	if exists {
		netAddress := new(net_address.NetAddress)
		err := netAddress.Set(a)
		if err == nil {
			env.NetAddress = netAddress
		}
	}

	pi, exists := os.LookupEnv("POLL_INTERVAL")
	if exists {
		i, err := strconv.Atoi(pi)
		if err == nil {
			env.PollInterval = &i
		}
	}

	ri, exists := os.LookupEnv("REPORT_INTERVAL")
	if exists {
		i, err := strconv.Atoi(ri)
		if err == nil {
			env.ReportInterval = &i
		}
	}

	ale, exists := os.LookupEnv("AGENT_LOG_ENABLE")
	if exists {
		i, err := strconv.Atoi(ale)
		if err == nil {
			b := i == 1
			env.AgentLogEnable = &b
		}
	}
}

func GetEnv() *base.Config {
	return env
}
