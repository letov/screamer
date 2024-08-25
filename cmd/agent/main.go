package main

import (
	"screamer/internal/collector"
	"screamer/internal/config"
	"screamer/internal/loop-agent"
)

func init() {
	config.InitAgent()
	collector.Init()
}

func main() {
	loop_agent.Run()
}
