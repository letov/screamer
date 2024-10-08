package config

import (
	"os"
	net_address "screamer/internal/common/netaddress"
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

	return pre
}
