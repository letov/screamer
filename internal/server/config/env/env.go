package env

import (
	"os"
	net_address "screamer/internal/common/net-address"
	"screamer/internal/server/config/base"
	"strconv"
)

var env *base.Config

func Init() {
	env = &base.Config{
		NetAddress:      nil,
		StoreInterval:   nil,
		FileStoragePath: nil,
		Restore:         nil,
		ServerLogEnable: nil,
	}

	a, exists := os.LookupEnv("ADDRESS")
	if exists {
		netAddress := new(net_address.NetAddress)
		err := netAddress.Set(a)
		if err == nil {
			env.NetAddress = netAddress
		}
	}

	si, exists := os.LookupEnv("STORE_INTERVAL")
	if exists {
		i, err := strconv.Atoi(si)
		if err == nil {
			env.StoreInterval = &i
		}
	}

	fsp, exists := os.LookupEnv("FILE_STORAGE_PATH")
	if exists {
		env.FileStoragePath = &fsp
	}

	r, exists := os.LookupEnv("RESTORE")
	if exists {
		i, err := strconv.Atoi(r)
		if err == nil {
			b := i == 1
			env.Restore = &b
		}
	}

	sle, exists := os.LookupEnv("SERVER_LOG_ENABLE")
	if exists {
		i, err := strconv.Atoi(sle)
		if err == nil {
			b := i == 1
			env.Restore = &b
		}
	}
}

func GetEnv() *base.Config {
	return env
}
