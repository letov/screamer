package env

import (
	"github.com/caarlos0/env"
	"log"
)

var envS EnvS

type EnvS struct {
	Address string `env:"ADDRESS"`
}

func InitServer() {
	err := env.Parse(&envS)
	if err != nil {
		log.Fatal(err)
	}
}

func GetEnvS() *EnvS {
	return &envS
}
