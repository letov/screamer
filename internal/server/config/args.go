package config

import (
	"flag"
	net_address "screamer/internal/common/net-address"
)

func newArgs() preConfig {
	netAddress := new(net_address.NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")

	pre := preConfig{
		NetAddress:      netAddress,
		StoreInterval:   flag.Int("i", 300, "StoreInterval desc"),
		FileStoragePath: flag.String("f", "data", "FileStoragePath desc"),
		Restore:         flag.Bool("r", true, "Restore desc"),
		ServerLogEnable: flag.Bool("l", true, "ServerLogEnable desc"),
	}

	set := newSetConfig()

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			set.NetAddress = true
			break
		case "i":
			set.StoreInterval = true
			break
		case "f":
			set.FileStoragePath = true
			break
		case "r":
			set.Restore = true
			break
		case "l":
			set.ServerLogEnable = true
			break
		}
	})

	flag.Parse()

	if !set.NetAddress {
		pre.NetAddress = nil
	}
	if !set.StoreInterval {
		pre.StoreInterval = nil
	}
	if !set.FileStoragePath {
		pre.FileStoragePath = nil
	}
	if !set.Restore {
		pre.Restore = nil
	}
	if !set.ServerLogEnable {
		pre.ServerLogEnable = nil
	}

	return pre
}
