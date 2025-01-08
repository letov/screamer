package config

import (
	"os"
	"screamer/internal/common/netaddress"
	"strconv"

	"github.com/joho/godotenv"
)

func newDotenv() preConfig {
	if os.Getenv("IS_TEST_ENV") == "true" {
		_ = godotenv.Load("../../.env.agent.test")
	} else {
		err := godotenv.Load(".env.agent.local")
		if err != nil {
			_ = godotenv.Load(".env.agent")
		}
	}

	netAddress := new(netaddress.NetAddress)
	_ = netAddress.Set(*getEnv("ADDRESS", ""))

	return preConfig{
		NetAddress:     netAddress,
		PollInterval:   getEnvInt("POLL_INTERVAL", 0),
		ReportInterval: getEnvInt("REPORT_INTERVAL", 0),
		Key:            getEnv("KEY", ""),
		RateLimit:      getEnvInt("RATE_LIMIT", 0),
		CryptoKey:      getEnv("CRYPTO_KEY", ""),
		Host:           getEnv("HOST", ""),
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
