package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var config *Config

type Config struct {
	Port string
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
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
