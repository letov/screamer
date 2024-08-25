package env

import (
	"os"
	"strconv"
)

var envS EnvS
var envSSet EnvSSet

type EnvS struct {
	Address         string `env:"ADDRESS"`
	ServerLogEnable bool   `env:"SERVER_LOG_ENABLE"`
	StoreInterval   int    `env:"STORE_INTERVAL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	Restore         bool   `env:"RESTORE"`
}

type EnvSSet struct {
	Address         bool
	ServerLogEnable bool
	StoreInterval   bool
	FileStoragePath bool
	Restore         bool
}

func InitServer() {
	address, exists := os.LookupEnv("ADDRESS")
	if exists {
		envS.Address = address
		envSSet.Address = true
	}

	serverLogEnable, exists := os.LookupEnv("SERVER_LOG_ENABLE")
	if exists {
		envS.ServerLogEnable = serverLogEnable == "true"
		envSSet.Address = true
	}

	storeInterval, exists := os.LookupEnv("STORE_INTERVAL")
	if exists {
		v, err := strconv.Atoi(storeInterval)
		if err != nil {
			envS.StoreInterval = v
			envSSet.Address = true
		}
	}

	fileStoragePath, exists := os.LookupEnv("FILE_STORAGE_PATH")
	if exists {
		envS.FileStoragePath = fileStoragePath
		envSSet.Address = true
	}

	restore, exists := os.LookupEnv("RESTORE")
	if exists {
		envS.Restore = restore == "true"
		envSSet.Address = true
	}
}

func GetEnvS() *EnvS {
	return &envS
}

func GetEnvSSet() *EnvSSet {
	return &envSSet
}
