package main

import (
	"screamer/internal/config"
	"screamer/internal/grab"
	"screamer/internal/storage"
)

func init() {
	config.Init()
	storage.Init()
	grab.Init()
}

func main() {

}
