package config

import (
	"os"
	net_address "screamer/internal/common/helpers/netaddress"
	"strconv"

	"github.com/joho/godotenv"
)

func newDotenv() preConfig {
	if os.Getenv("IS_TEST_ENV") == "true" {
		_ = godotenv.Load("../../../.env.server.test")
	} else {
		err := godotenv.Load(".env.server.local")
		if err != nil {
			_ = godotenv.Load(".env.server")
		}
	}

	netAddress := new(net_address.NetAddress)
	_ = netAddress.Set(*getEnv("ADDRESS", ""))

	netAddressGrpc := new(net_address.NetAddress)
	_ = netAddressGrpc.Set(*getEnv("ADDRESS_GRPC", ""))

	r := *getEnv("RESTORE", "false")
	br, err := strconv.ParseBool(r)
	if err != nil {
		br = false
	}

	return preConfig{
		NetAddress:      netAddress,
		DBAddress:       getEnv("DATABASE_DSN", ""),
		StoreInterval:   getEnvInt("STORE_INTERVAL", 0),
		FileStoragePath: getEnv("FILE_STORAGE_PATH", ""),
		Restore:         &br,
		Key:             getEnv("KEY", ""),
		CryptoKey:       getEnv("CRYPTO_KEY", ""),
		TrustedSubnet:   getEnv("TRUSTED_SUBNET", ""),
		NetAddressGrpc:  netAddressGrpc,
	}
}

func getEnvInt(key string, def int) *int {
	v, e := strconv.Atoi(*getEnv(key, strconv.Itoa(def)))
	if e != nil {
		return &def
	} else {
		return &v
	}
}

func getEnv(key string, def string) *string {
	if value, exists := os.LookupEnv(key); exists {
		return &value
	}
	return &def
}
