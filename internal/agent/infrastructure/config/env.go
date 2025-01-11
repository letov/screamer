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
		if err := netAddress.Set(a); err == nil {
			pre.NetAddress = netAddress
		}
	}

	pi, exists := os.LookupEnv("POLL_INTERVAL")
	if exists {
		if i, err := strconv.Atoi(pi); err == nil {
			pre.PollInterval = &i
		}
	}

	ri, exists := os.LookupEnv("REPORT_INTERVAL")
	if exists {
		if i, err := strconv.Atoi(ri); err == nil {
			pre.ReportInterval = &i
		}
	}

	k, exists := os.LookupEnv("KEY")
	if exists {
		pre.Key = &k
	}

	rl, exists := os.LookupEnv("RATE_LIMIT")
	if exists {
		if i, err := strconv.Atoi(rl); err == nil {
			pre.RateLimit = &i
		}
	}

	ck, exists := os.LookupEnv("CRYPTO_KEY")
	if exists {
		pre.CryptoKey = &ck
	}

	h, exists := os.LookupEnv("HOST")
	if exists {
		pre.Host = &h
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
