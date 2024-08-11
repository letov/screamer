package main

import (
	"screamer/internal/collector"
	"screamer/internal/config"
	"screamer/internal/loop"
)

func init() {
	config.InitAgent()
	collector.Init()
}

func main() {
	loop.Run()
}
