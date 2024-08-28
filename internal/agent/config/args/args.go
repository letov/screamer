package args

import (
	"flag"
	"screamer/internal/agent/config/base"
	net_address "screamer/internal/common/net-address"
)

type ArgsSet struct {
	NetAddress     bool
	PollInterval   bool
	ReportInterval bool
	AgentLogEnable bool
}

var args *base.Config
var argsSet ArgsSet

func Init() {
	netAddress := new(net_address.NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")

	args = &base.Config{
		NetAddress:     netAddress,
		PollInterval:   flag.Int("p", 0, "PollInterval desc"),
		ReportInterval: flag.Int("r", 0, "ReportInterval desc"),
		AgentLogEnable: flag.Bool("l", true, "AgentLogEnable desc"),
	}

	argsSet = ArgsSet{
		NetAddress:     false,
		PollInterval:   false,
		ReportInterval: false,
		AgentLogEnable: false,
	}

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			argsSet.NetAddress = true
			break
		case "p":
			argsSet.PollInterval = true
			break
		case "r":
			argsSet.ReportInterval = true
			break
		case "l":
			argsSet.AgentLogEnable = true
			break
		}
	})

	flag.Parse()
}

func GetArgs() *base.Config {
	res := args

	if !argsSet.NetAddress {
		res.NetAddress = nil
	}
	if !argsSet.PollInterval {
		res.PollInterval = nil
	}
	if !argsSet.ReportInterval {
		res.ReportInterval = nil
	}
	if !argsSet.AgentLogEnable {
		res.AgentLogEnable = nil
	}

	return res
}
