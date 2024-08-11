package dotenv

import (
	"github.com/joho/godotenv"
	"log"
)

var dotenvS *DotenvS

type DotenvS struct {
	Address         string
	ServerLogEnable bool
}

func InitServer() {
	err := godotenv.Load(".env.server")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dotenvS = &DotenvS{
		Address:         getEnv("ADDRESS", "localhost:8080"),
		ServerLogEnable: getEnvInt("SERVER_LOG_ENABLE", 1) == 1,
	}
}

func GetDotenvS() *DotenvS {
	return dotenvS
}
