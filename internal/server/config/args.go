package config

import (
	"flag"
	net_address "screamer/internal/common/netaddress"
)

func newArgs() preConfig {
	netAddress := new(net_address.NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")

	pre := preConfig{
		NetAddress:      netAddress,
		StoreInterval:   flag.Int("i", 300, "StoreInterval desc"),
		FileStoragePath: flag.String("f", "/tmp/backup_file", "FileStoragePath desc"),
		Restore:         flag.Bool("r", true, "Restore desc"),
		ServerLogEnable: flag.Bool("l", true, "ServerLogEnable desc"),
	}

	set := newSetConfig()

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			set.NetAddress = true
		case "i":
			set.StoreInterval = true
		case "f":
			set.FileStoragePath = true
		case "r":
			set.Restore = true
		case "l":
			set.ServerLogEnable = true
		}
	})

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
