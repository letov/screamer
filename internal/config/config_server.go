package config

import (
	"screamer/internal/config/args"
	"screamer/internal/config/dotenv"
	"screamer/internal/config/env"
)

var configS *ConfigS

type ConfigS struct {
	Port            string
	ServerLogEnable bool
}

type EnvSrcS struct {
	ArgsS   *args.ArgsS
	EnvS    *env.EnvS
	DotenvS *dotenv.DotenvS
}

func InitServer() {
	args.InitServer()
	env.InitServer()
	dotenv.InitServer()

	envSrcS := EnvSrcS{
		ArgsS:   args.GetArgsS(),
		EnvS:    env.GetEnvS(),
		DotenvS: dotenv.GetDotenvS(),
	}

	configS = &ConfigS{
		Port:            getPort(&envSrcS),
		ServerLogEnable: envSrcS.DotenvS.ServerLogEnable,
	}
}

func GetConfigS() *ConfigS {
	return configS
}

func getPort(envSrcS *EnvSrcS) string {
	var serverURL string
	if envSrcS.EnvS.Address != "" {
		serverURL = envSrcS.EnvS.Address
	} else if envSrcS.ArgsS.NetAddress.Host != "" {
		serverURL = envSrcS.ArgsS.NetAddress.String()
	} else {
		serverURL = envSrcS.DotenvS.Address
	}
	return getPortFromUrl(serverURL)
}
