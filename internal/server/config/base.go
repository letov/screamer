package config

import net_address "screamer/internal/common/netaddress"

type Config struct {
	NetAddress      net_address.NetAddress
	DBAddress       string
	StoreInterval   int
	FileStoragePath string
	Restore         bool
}

type preConfig struct {
	NetAddress      *net_address.NetAddress
	DBAddress       *string
	StoreInterval   *int
	FileStoragePath *string
	Restore         *bool
}

type setConfig struct {
	NetAddress      bool
	DBAddress       bool
	StoreInterval   bool
	FileStoragePath bool
	Restore         bool
}

func newPreConfig() preConfig {
	return preConfig{
		NetAddress:      nil,
		DBAddress:       nil,
		StoreInterval:   nil,
		FileStoragePath: nil,
		Restore:         nil,
	}
}

func newSetConfig() setConfig {
	return setConfig{
		NetAddress:      false,
		DBAddress:       false,
		StoreInterval:   false,
		FileStoragePath: false,
		Restore:         false,
	}
}
