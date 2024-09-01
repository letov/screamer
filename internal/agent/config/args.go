package config

import (
	"flag"
	"fmt"
	net_address "screamer/internal/common/net-address"
)

func newArgs() preConfig {
	netAddress := new(net_address.NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")

	pre := preConfig{
		NetAddress:     netAddress,
		PollInterval:   flag.Int("p", 0, "PollInterval desc"),
		ReportInterval: flag.Int("r", 0, "ReportInterval desc"),
		AgentLogEnable: flag.Bool("l", true, "AgentLogEnable desc"),
	}

	set := newSetConfig()

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			set.NetAddress = true
			break
		case "p":
			set.PollInterval = true
			break
		case "r":
			set.ReportInterval = true
			break
		case "l":
			set.AgentLogEnable = true
			break
		}
	})

	if !set.NetAddress {
		pre.NetAddress = nil
	}
	if !set.PollInterval {
		pre.PollInterval = nil
	}
	if !set.ReportInterval {
		pre.ReportInterval = nil
	}
	if !set.AgentLogEnable {
		pre.AgentLogEnable = nil
	}

	fmt.Println(pre)

	return pre
}
