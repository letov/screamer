package args

import (
	"flag"
)

type ArgsA struct {
	NetAddress     *NetAddress
	PollInterval   *int
	ReportInterval *int
}

var argsA *ArgsA

func InitAgent() {
	netAddress := new(NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")
	argsA = &ArgsA{
		NetAddress:     netAddress,
		PollInterval:   flag.Int("p", 0, "pollInterval desc"),
		ReportInterval: flag.Int("r", 0, "reportInterval desc"),
	}
	flag.Parse()
}

func GetArgsA() *ArgsA {
	return argsA
}
