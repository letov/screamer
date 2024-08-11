package main

import (
	"screamer/internal/args"
	"screamer/internal/collector"
	"screamer/internal/config"
	"screamer/internal/loop"
)

func init() {
	args.InitAgent()
	config.InitAgent()
	collector.Init()
}

func main() {
	loop.Run()
}
