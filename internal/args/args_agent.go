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
	flag.Var(netAddress, "a", "Server address host:port")
	argsAgent = &ArgsAgent{
		NetAddress:     netAddress,
		PollInterval:   flag.Int("p", UnsetIntValue, "pollInterval desc"),
		ReportInterval: flag.Int("r", UnsetIntValue, "reportInterval desc"),
	}
	flag.Parse()
}

func GetArgsAgent() *ArgsAgent {
	return argsAgent
}
