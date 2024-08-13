package main

import (
	"screamer/internal/config"
	"screamer/internal/grab"
	"screamer/internal/logger"
	"screamer/internal/storage"
)

func init() {
	config.InitServer()
	storage.Init()
	logger.Init()
	grab.Init()
}

func main() {

}
