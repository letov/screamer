package main

import (
	"screamer/internal/config"
	"screamer/internal/grab"
	"screamer/internal/logger"
	"screamer/internal/loop-server"
	"screamer/internal/storage"
)

func init() {
	config.InitServer()
	storage.Init()
	logger.Init()
	go grab.Init()
}

func main() {
	defer serverStopped()
	loop_server.Run()
}

func serverStopped() {

}
