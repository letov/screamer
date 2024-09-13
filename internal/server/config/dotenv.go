package config

import (
	"github.com/joho/godotenv"
	"os"
	net_address "screamer/internal/common/netaddress"
	"strconv"
)

func newDotenv() preConfig {
	_ = godotenv.Load(".env.server")

	netAddress := new(net_address.NetAddress)
	_ = netAddress.Set(*getEnv("ADDRESS", ""))

	r := *getEnvInt("RESTORE", 0) == 1

	return preConfig{
		NetAddress:      netAddress,
		DBAddress:       getEnv("DATABASE_DSN", ""),
		StoreInterval:   getEnvInt("STORE_INTERVAL", 0),
		FileStoragePath: getEnv("FILE_STORAGE_PATH", ""),
		Restore:         &r,
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
