package main

import (
	"screamer/internal/args"
	"screamer/internal/config"
	"screamer/internal/grab"
	"screamer/internal/storage"
)

func init() {
	args.InitServer()
	config.InitServer()
	storage.Init()
	grab.Init()
}

func main() {

}
