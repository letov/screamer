package config

import (
	"fmt"
	"screamer/internal/args"
)

var configServer *ConfigServer

type ConfigServer struct {
	Port            string
	ServerLogEnable bool
}

func InitServer() {
	Init()
	configServer = newConfigServer()
}

func GetConfigServer() *ConfigServer {
	return configServer
}

func newConfigServer() *ConfigServer {
	a := args.GetArgsServer()

	var port string
	if a.NetAddress.Port == UnsetIntValue {
		port = getEnv("PORT", "8080")
	} else {
		port = fmt.Sprintf("%d", a.NetAddress.Port)
	}

	return &ConfigServer{
		Port:            port,
		ServerLogEnable: getEnvInt("SERVER_LOG_ENABLE", 1) == 1,
	}
}
