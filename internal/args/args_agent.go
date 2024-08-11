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
	flag.Var(netAddress, "a", "Server address host:port")
	argsAgent = &ArgsAgent{
		NetAddress:     netAddress,
		PollInterval:   flag.Int("p", 0, "pollInterval desc"),
		ReportInterval: flag.Int("r", 0, "reportInterval desc"),
	}
	flag.Parse()
}

func GetArgsAgent() *ArgsAgent {
	return argsAgent
}
