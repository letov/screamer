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
	StoreInterval   int
	FileStoragePath string
	Restore         bool
}

type EnvSrcS struct {
	ArgsS    *args.ArgsS
	ArgsSSet *args.ArgsSSet
	EnvS     *env.EnvS
	EnvSSet  *env.EnvSSet
	DotenvS  *dotenv.DotenvS
}

func InitServer() {
	args.InitServer()
	env.InitServer()
	dotenv.InitServer()

	envSrcS := EnvSrcS{
		ArgsS:    args.GetArgsS(),
		ArgsSSet: args.GetArgsSSet(),
		EnvS:     env.GetEnvS(),
		EnvSSet:  env.GetEnvSSet(),
		DotenvS:  dotenv.GetDotenvS(),
	}

	configS = &ConfigS{
		Port:            getPort(&envSrcS),
		ServerLogEnable: getServerLogEnable(&envSrcS),
		StoreInterval:   getStoreInterval(&envSrcS),
		FileStoragePath: getFileStoragePath(&envSrcS),
		Restore:         getRestore(&envSrcS),
	}
}

func GetConfigS() *ConfigS {
	return configS
}

func getPort(e *EnvSrcS) string {
	var serverURL string
	if e.EnvSSet.Address {
		serverURL = e.EnvS.Address
	} else if e.ArgsSSet.NetAddress {
		serverURL = e.ArgsS.NetAddress.String()
	} else {
		serverURL = e.DotenvS.Address
	}
	return getPortFromURL(serverURL)
}

func getRestore(e *EnvSrcS) bool {
	if e.EnvSSet.Restore {
		return e.EnvS.Restore
	} else if e.ArgsSSet.Restore {
		return *e.ArgsS.Restore
	}

	return e.DotenvS.Restore
}

func getFileStoragePath(e *EnvSrcS) string {
	if e.EnvSSet.FileStoragePath {
		return e.EnvS.FileStoragePath
	} else if e.ArgsSSet.FileStoragePath {
		return *e.ArgsS.FileStoragePath
	}

	return e.DotenvS.FileStoragePath
}

func getStoreInterval(e *EnvSrcS) int {
	if e.EnvSSet.StoreInterval {
		return e.EnvS.StoreInterval
	} else if e.ArgsSSet.StoreInterval {
		return *e.ArgsS.StoreInterval
	}

	return e.DotenvS.StoreInterval
}

func getServerLogEnable(e *EnvSrcS) bool {
	if e.EnvSSet.ServerLogEnable {
		return e.EnvS.ServerLogEnable
	} else if e.ArgsSSet.ServerLogEnable {
		return *e.ArgsS.ServerLogEnable
	}

	return e.DotenvS.ServerLogEnable
}
