package config

import (
	"github.com/joho/godotenv"
	"os"
	net_address "screamer/internal/common/netaddress"
	"strconv"
)

func newDotenv() preConfig {
	_ = godotenv.Load(".env.agent")

	netAddress := new(net_address.NetAddress)
	_ = netAddress.Set(*getEnv("ADDRESS", "localhost:8080"))

	r := *getEnvInt("RESTORE", 1) == 1

	return preConfig{
		NetAddress:      netAddress,
		DBAddress:       getEnv("FILE_STORAGE_PATH", "postgres://my_user:my_pass@localhost:25432/my_db"),
		StoreInterval:   getEnvInt("STORE_INTERVAL", 300),
		FileStoragePath: getEnv("FILE_STORAGE_PATH", "/tmp/backup_file"),
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
