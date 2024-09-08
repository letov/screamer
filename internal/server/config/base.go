package config

import net_address "screamer/internal/common/netaddress"

type Config struct {
	NetAddress      net_address.NetAddress
	StoreInterval   int
	FileStoragePath string
	Restore         bool
	ServerLogEnable bool
}

type preConfig struct {
	NetAddress      *net_address.NetAddress
	StoreInterval   *int
	FileStoragePath *string
	Restore         *bool
	ServerLogEnable *bool
}

type setConfig struct {
	NetAddress      bool
	StoreInterval   bool
	FileStoragePath bool
	Restore         bool
	ServerLogEnable bool
}

func newPreConfig() preConfig {
	return preConfig{
		NetAddress:      nil,
		StoreInterval:   nil,
		FileStoragePath: nil,
		Restore:         nil,
		ServerLogEnable: nil,
	}
}

func newSetConfig() setConfig {
	return setConfig{
		NetAddress:      false,
		StoreInterval:   false,
		FileStoragePath: false,
		Restore:         false,
		ServerLogEnable: false,
	}
}
