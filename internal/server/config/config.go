package config

import (
	net_address "screamer/internal/common/net-address"
	"screamer/internal/server/config/args"
	"screamer/internal/server/config/base"
	"screamer/internal/server/config/dotenv"
	"screamer/internal/server/config/env"
)

var config *base.Config

type EnvSrc struct {
	Args   *base.Config
	Env    *base.Config
	Dotenv *base.Config
}

func Init() {
	args.Init()
	env.Init()
	dotenv.Init()

	envSrc := EnvSrc{
		Args:   args.GetArgs(),
		Env:    env.GetEnv(),
		Dotenv: dotenv.GetDotenv(),
	}

	config = &base.Config{
		NetAddress:      getNetAddress(envSrc),
		ServerLogEnable: getServerLogEnable(envSrc),
		StoreInterval:   getStoreInterval(envSrc),
		FileStoragePath: getFileStoragePath(envSrc),
		Restore:         getRestore(envSrc),
	}
}

func GetConfig() *base.Config {
	return config
}

func getNetAddress(envSrc EnvSrc) *net_address.NetAddress {
	if envSrc.Env.NetAddress != nil {
		return envSrc.Env.NetAddress
	} else if envSrc.Args.NetAddress != nil {
		return envSrc.Args.NetAddress
	}
	return envSrc.Dotenv.NetAddress
}

func getRestore(envSrc EnvSrc) *bool {
	if envSrc.Env.Restore != nil {
		return envSrc.Env.Restore
	} else if envSrc.Args.Restore != nil {
		return envSrc.Args.Restore
	}
	return envSrc.Dotenv.Restore
}

func getFileStoragePath(envSrc EnvSrc) *string {
	if envSrc.Env.FileStoragePath != nil {
		return envSrc.Env.FileStoragePath
	} else if envSrc.Args.FileStoragePath != nil {
		return envSrc.Args.FileStoragePath
	}
	return envSrc.Dotenv.FileStoragePath
}

func getStoreInterval(envSrc EnvSrc) *int {
	if envSrc.Env.StoreInterval != nil {
		return envSrc.Env.StoreInterval
	} else if envSrc.Args.StoreInterval != nil {
		return envSrc.Args.StoreInterval
	}

	return envSrc.Dotenv.StoreInterval
}

func getServerLogEnable(envSrc EnvSrc) *bool {
	if envSrc.Env.ServerLogEnable != nil {
		return envSrc.Env.ServerLogEnable
	} else if envSrc.Args.ServerLogEnable != nil {
		return envSrc.Args.ServerLogEnable
	}

	return envSrc.Dotenv.ServerLogEnable
}
