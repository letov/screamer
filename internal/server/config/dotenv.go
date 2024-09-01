package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	net_address "screamer/internal/common/netaddress"
	"strconv"
)

func newDotenv() preConfig {
	if err := godotenv.Load(".env.server"); err != nil {
		log.Println("Error loading .env file")
	}

	netAddress := new(net_address.NetAddress)
	_ = netAddress.Set(*getEnv("ADDRESS", "localhost:8080"))

	sle := *getEnvInt("SERVER_LOG_ENABLE", 1) == 1
	r := *getEnvInt("RESTORE", 1) == 1

	return preConfig{
		NetAddress:      netAddress,
		StoreInterval:   getEnvInt("STORE_INTERVAL", 300),
		FileStoragePath: getEnv("FILE_STORAGE_PATH", "/tmp/backup_file"),
		Restore:         &r,
		ServerLogEnable: &sle,
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
