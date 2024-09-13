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
