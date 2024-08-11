package args

import (
	"flag"
)

type ArgsAgent struct {
	NetAddress     *NetAddress
	PollInterval   *int
	ReportInterval *int
}

var argsAgent *ArgsAgent

func InitAgent() {
	netAddress := new(NetAddress)
	var _ = flag.Value(netAddress)
	flag.Var(netAddress, "addr", "Server address host:port")
	flag.Parse()
	argsAgent = &ArgsAgent{
		NetAddress:     netAddress,
		PollInterval:   flag.Int("p", UnsetIntValue, "pollInterval desc"),
		ReportInterval: flag.Int("r", UnsetIntValue, "reportInterval desc"),
	}
}

func GetArgsAgent() *ArgsAgent {
	return argsAgent
}
