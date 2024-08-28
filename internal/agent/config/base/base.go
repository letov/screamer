package base

import "screamer/internal/common/net-address"

type Config struct {
	NetAddress     *net_address.NetAddress
	PollInterval   *int
	ReportInterval *int
	AgentLogEnable *bool
}
