package config

import (
	"flag"
	"os"
	net_address "screamer/internal/common/helpers/netaddress"
)

func newArgs() preConfig {
	if os.Getenv("IS_TEST_ENV") == "true" {
		return preConfig{}
	}

	netAddress := new(net_address.NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")

	netAddressGrpc := new(net_address.NetAddress)
	flag.Var(netAddressGrpc, "ag", "Grpc server address host:port")

	pre := preConfig{
		NetAddress:      netAddress,
		DBAddress:       flag.String("d", "", "DBAddress desc"),
		StoreInterval:   flag.Int("i", 0, "StoreInterval desc"),
		FileStoragePath: flag.String("f", "", "FileStoragePath desc"),
		Restore:         flag.Bool("r", false, "Restore desc"),
		Key:             flag.String("k", "", "Key desc"),
		CryptoKey:       flag.String("crypto-key", "", "CryptoKey desc"),
		TrustedSubnet:   flag.String("t", "", "TrustedSubnet desc"),
		NetAddressGrpc:  netAddressGrpc,
	}

	set := newSetConfig()

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			set.NetAddress = true
		case "d":
			set.DBAddress = true
		case "i":
			set.StoreInterval = true
		case "f":
			set.FileStoragePath = true
		case "r":
			set.Restore = true
		case "k":
			set.Key = true
		case "crypto-key":
			set.CryptoKey = true
		case "t":
			set.CryptoKey = true
		case "ag":
			set.NetAddressGrpc = true
		}
	})

	if !set.NetAddress {
		pre.NetAddress = nil
	}
	if !set.DBAddress {
		pre.DBAddress = nil
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
	if !set.Key {
		pre.Key = nil
	}
	if !set.CryptoKey {
		pre.CryptoKey = nil
	}
	if !set.TrustedSubnet {
		pre.TrustedSubnet = nil
	}
	if !set.NetAddressGrpc {
		pre.NetAddressGrpc = nil
	}

	return pre
}
