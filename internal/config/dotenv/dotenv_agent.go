package dotenv

import (
	"github.com/joho/godotenv"
	"log"
)

var dotenvA *DotenvA

type DotenvA struct {
	PollInterval   int
	ReportInterval int
	Address        string
	AgentLogEnable bool
}

func InitAgent() {
	err := godotenv.Load(".env.agent")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dotenvA = &DotenvA{
		PollInterval:   getEnvInt("POLL_INTERVAL", 2),
		ReportInterval: getEnvInt("REPORT_INTERVAL", 10),
		Address:        getEnv("ADDRESS", "localhost:8080"),
		AgentLogEnable: getEnvInt("AGENT_LOG_ENABLE", 1) == 1,
	}
}

func GetDotenvA() *DotenvA {
	return dotenvA
}
