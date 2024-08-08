package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var config *Config

type Config struct {
	Port            string
	PollInterval    int
	ReportInterval  int
	ServerUrl       string
	AgentLogEnable  bool
	ServerLogEnable bool
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config = newConfig()
}

func GetConfig() *Config {
	return config
}

func newConfig() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		PollInterval:    getEnvInt("POLL_INTERVAL", 2),
		ReportInterval:  getEnvInt("REPORT_INTERVAL", 10),
		ServerUrl:       getEnv("SERVER_URL", "http://localhost:8080"),
		AgentLogEnable:  getEnvInt("AGENT_LOG_ENABLE", 1) == 1,
		ServerLogEnable: getEnvInt("SERVER_LOG_ENABLE", 1) == 1,
	}
}

func getEnvInt(key string, def int) int {
	v, e := strconv.Atoi(getEnv(key, strconv.Itoa(def)))
	if e != nil {
		return def
	} else {
		return v
	}
}

func getEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
