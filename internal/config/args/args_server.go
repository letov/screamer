package args

import (
	"flag"
)

type ArgsS struct {
	NetAddress      *NetAddress
	ServerLogEnable *bool
	StoreInterval   *int
	FileStoragePath *string
	Restore         *bool
}

type ArgsSSet struct {
	NetAddress      bool
	ServerLogEnable bool
	StoreInterval   bool
	FileStoragePath bool
	Restore         bool
}

var argsS *ArgsS
var argsSSet ArgsSSet

func InitServer() {
	netAddress := new(NetAddress)
	flag.Var(netAddress, "a", "Server address host:port")

	argsS = &ArgsS{
		NetAddress:      netAddress,
		ServerLogEnable: flag.Bool("l", true, "ServerLogEnable desc"),
		StoreInterval:   flag.Int("i", 300, "StoreInterval desc"),
		FileStoragePath: flag.String("f", "./store/storage", "FileStoragePath desc"),
		Restore:         flag.Bool("r", true, "Restore desc"),
	}

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "a":
			argsSSet.NetAddress = true
			break
		case "l":
			argsSSet.ServerLogEnable = true
			break
		case "i":
			argsSSet.StoreInterval = true
			break
		case "f":
			argsSSet.FileStoragePath = true
			break
		case "r":
			argsSSet.Restore = true
			break
		}
	})

	flag.Parse()
}

func GetArgsS() *ArgsS {
	return argsS
}

func GetArgsSSet() *ArgsSSet {
	return &argsSSet
}
