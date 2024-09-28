package config

import net_address "screamer/internal/common/netaddress"

type Config struct {
	NetAddress      net_address.NetAddress
	DBAddress       string
	StoreInterval   int
	FileStoragePath string
	Restore         bool
	Key             string
}

type preConfig struct {
	NetAddress      *net_address.NetAddress
	DBAddress       *string
	StoreInterval   *int
	FileStoragePath *string
	Restore         *bool
	Key             *string
}

type setConfig struct {
	NetAddress      bool
	DBAddress       bool
	StoreInterval   bool
	FileStoragePath bool
	Restore         bool
	Key             bool
}

func newPreConfig() preConfig {
	return preConfig{
		NetAddress:      nil,
		DBAddress:       nil,
		StoreInterval:   nil,
		FileStoragePath: nil,
		Restore:         nil,
		Key:             nil,
	}
}

func newSetConfig() setConfig {
	return setConfig{
		NetAddress:      false,
		DBAddress:       false,
		StoreInterval:   false,
		FileStoragePath: false,
		Restore:         false,
		Key:             false,
	}
}
