package config

import (
	"os"
	net_address "screamer/internal/common/netaddress"
	"strconv"
)

func newEnv() preConfig {
	pre := newPreConfig()

	a, exists := os.LookupEnv("ADDRESS")
	if exists {
		netAddress := new(net_address.NetAddress)
		if err := netAddress.Set(a); err == nil {
			pre.NetAddress = netAddress
		}
	}

	pi, exists := os.LookupEnv("POLL_INTERVAL")
	if exists {
		if i, err := strconv.Atoi(pi); err == nil {
			pre.PollInterval = &i
		}
	}

	ri, exists := os.LookupEnv("REPORT_INTERVAL")
	if exists {
		if i, err := strconv.Atoi(ri); err == nil {
			pre.ReportInterval = &i
		}
	}

	return pre
}
