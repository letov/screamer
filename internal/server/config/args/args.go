package args

import (
	"flag"
	net_address "screamer/internal/common/net-address"
	"screamer/internal/server/config/base"
)

type ArgsSet struct {
	NetAddress      bool
	ServerLogEnable bool
	StoreInterval   bool
	FileStoragePath bool
	Restore         bool
}

var args *base.Config
var argsSet ArgsSet

func Init() {
	netAddress := new(net_address.NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")

	args = &base.Config{
		NetAddress:      netAddress,
		StoreInterval:   flag.Int("i", 300, "StoreInterval desc"),
		FileStoragePath: flag.String("f", "data", "FileStoragePath desc"),
		Restore:         flag.Bool("r", true, "Restore desc"),
		ServerLogEnable: flag.Bool("l", true, "ServerLogEnable desc"),
	}

	argsSet = ArgsSet{
		NetAddress:      false,
		StoreInterval:   false,
		FileStoragePath: false,
		Restore:         false,
		ServerLogEnable: false,
	}

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			argsSet.NetAddress = true
			break
		case "i":
			argsSet.StoreInterval = true
			break
		case "f":
			argsSet.FileStoragePath = true
			break
		case "r":
			argsSet.Restore = true
			break
		case "l":
			argsSet.ServerLogEnable = true
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
	if !argsSet.StoreInterval {
		res.StoreInterval = nil
	}
	if !argsSet.FileStoragePath {
		res.FileStoragePath = nil
	}
	if !argsSet.Restore {
		res.Restore = nil
	}
	if !argsSet.ServerLogEnable {
		res.ServerLogEnable = nil
	}

	return res
}
