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
	_ = flag.Value(netAddress)
	flag.Var(netAddress, "addr", "Server address host:port")
	flag.Parse()
	argsServer = &ArgsServer{NetAddress: netAddress}
}

func GetArgsServer() *ArgsServer {
	return argsServer
}
