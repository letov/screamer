package main

import (
	"screamer/internal/args"
	"screamer/internal/config"
)

func init() {
	args.InitServer()
	config.InitServer()
	//storage.Init()
	//grab.Init()
}

func main() {

}
