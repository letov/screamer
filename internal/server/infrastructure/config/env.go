package config

import (
	"os"
	net_address "screamer/internal/common/helpers/netaddress"
	"strconv"
)

func newEnv() preConfig {
	pre := newPreConfig()

	a, exists := os.LookupEnv("ADDRESS")
	if exists {
		netAddress := new(net_address.NetAddress)
		err := netAddress.Set(a)
		if err == nil {
			pre.NetAddress = netAddress
		}
	}

	d, exists := os.LookupEnv("DATABASE_DSN")
	if exists {
		pre.DBAddress = &d
	}

	si, exists := os.LookupEnv("STORE_INTERVAL")
	if exists {
		i, err := strconv.Atoi(si)
		if err == nil {
			pre.StoreInterval = &i
		}
	}

	fsp, exists := os.LookupEnv("FILE_STORAGE_PATH")
	if exists {
		pre.FileStoragePath = &fsp
	}

	r, exists := os.LookupEnv("RESTORE")
	if exists {
		b, err := strconv.ParseBool(r)
		if err == nil {
			pre.Restore = &b
		}
	}

	k, exists := os.LookupEnv("KEY")
	if exists {
		pre.Key = &k
	}

	ck, exists := os.LookupEnv("CRYPTO_KEY")
	if exists {
		pre.CryptoKey = &ck
	}

	t, exists := os.LookupEnv("TRUSTED_SUBNET")
	if exists {
		pre.CryptoKey = &t
	}

	ag, exists := os.LookupEnv("ADDRESS_GRPC")
	if exists {
		netAddressGrpc := new(net_address.NetAddress)
		err := netAddressGrpc.Set(ag)
		if err == nil {
			pre.NetAddress = netAddressGrpc
		}
	}

	return pre
}
