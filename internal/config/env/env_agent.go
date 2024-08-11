package env

import (
	"github.com/caarlos0/env"
	"log"
)

var envA EnvA

type EnvA struct {
	PollInterval   int    `env:"POLL_INTERVAL"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	Address        string `env:"ADDRESS"`
}

func InitAgent() {
	err := env.Parse(&envA)
	if err != nil {
		log.Fatal(err)
	}
}

func GetEnvA() *EnvA {
	return &envA
}
