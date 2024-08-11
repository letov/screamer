package args

import (
	"flag"
)

type ArgsS struct {
	NetAddress *NetAddress
}

var argsS *ArgsS

func InitServer() {
	netAddress := new(NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")
	argsS = &ArgsS{NetAddress: netAddress}
	flag.Parse()
}

func GetArgsS() *ArgsS {
	return argsS
}
