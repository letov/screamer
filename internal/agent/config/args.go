package config

import (
	"flag"
	net_address "screamer/internal/common/netaddress"
)

func newArgs() preConfig {
	netAddress := new(net_address.NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")

	pre := preConfig{
		NetAddress:     netAddress,
		PollInterval:   flag.Int("p", 0, "PollInterval desc"),
		ReportInterval: flag.Int("r", 0, "ReportInterval desc"),
		Key:            flag.String("k", "", "Key desc"),
	}

	set := newSetConfig()

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			set.NetAddress = true
		case "p":
			set.PollInterval = true
		case "r":
			set.ReportInterval = true
		case "k":
			set.Key = true
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
	if !set.Key {
		pre.Key = nil
	}

	return pre
}
