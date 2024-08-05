package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewConfig() *Config {
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
