package base

import net_address "screamer/internal/common/net-address"

type Config struct {
	NetAddress      *net_address.NetAddress
	StoreInterval   *int
	FileStoragePath *string
	Restore         *bool
	ServerLogEnable *bool
}
