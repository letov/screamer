package config

import (
	"os"
	net_address "screamer/internal/common/net-address"
	"strconv"
)

func newEnv() preConfig {
	pre := newPreConfig()

	a, exists := os.LookupEnv("ADDRESS")
	if exists {
		netAddress := new(net_address.NetAddress)
		err := netAddress.Set(a)
		if err == nil {
			pre.NetAddress = netAddress
		}
	}

	pi, exists := os.LookupEnv("POLL_INTERVAL")
	if exists {
		i, err := strconv.Atoi(pi)
		if err == nil {
			pre.PollInterval = &i
		}
	}

	ri, exists := os.LookupEnv("REPORT_INTERVAL")
	if exists {
		i, err := strconv.Atoi(ri)
		if err == nil {
			pre.ReportInterval = &i
		}
	}

	ale, exists := os.LookupEnv("AGENT_LOG_ENABLE")
	if exists {
		i, err := strconv.Atoi(ale)
		if err == nil {
			b := i == 1
			pre.AgentLogEnable = &b
		}
	}

	return pre
}
