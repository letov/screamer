package dotenv

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"screamer/internal/agent/config/base"
	"screamer/internal/common/net-address"
	"strconv"
)

var dotenv *base.Config

func Init() {
	err := godotenv.Load(".env.agent")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	netAddress := new(net_address.NetAddress)
	_ = netAddress.Set(*getEnv("ADDRESS", "localhost:8080"))

	ale := *getEnvInt("AGENT_LOG_ENABLE", 1) == 1

	dotenv = &base.Config{
		NetAddress:     netAddress,
		PollInterval:   getEnvInt("POLL_INTERVAL", 2),
		ReportInterval: getEnvInt("REPORT_INTERVAL", 10),
		AgentLogEnable: &ale,
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

func GetDotenv() *base.Config {
	return dotenv
}
