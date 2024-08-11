package args

import (
	"flag"
)

type ArgsServer struct {
	NetAddress *NetAddress
}

var argsServer *ArgsServer

func InitServer() {
	netAddress := new(NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")
	argsServer = &ArgsServer{NetAddress: netAddress}
	flag.Parse()
}

func GetArgsServer() *ArgsServer {
	return argsServer
}
