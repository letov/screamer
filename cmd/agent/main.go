package main

import (
	"screamer/internal/collector"
	"screamer/internal/config"
	"screamer/internal/loop"
)

func init() {
	config.Init()
	collector.Init()
}

func main() {
	loop.Run()
}
