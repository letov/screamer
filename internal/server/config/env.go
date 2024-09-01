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
		i, err := strconv.Atoi(r)
		if err == nil {
			b := i == 1
			pre.Restore = &b
		}
	}

	sle, exists := os.LookupEnv("SERVER_LOG_ENABLE")
	if exists {
		i, err := strconv.Atoi(sle)
		if err == nil {
			b := i == 1
			pre.Restore = &b
		}
	}

	return pre
}
