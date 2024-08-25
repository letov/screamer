package loop_server

import (
	"math"
	"screamer/internal/backup"
	"screamer/internal/config"
	"time"
)

func Run() {
	c := config.GetConfigS()
	var i = 0
	for {
		if c.Restore && c.StoreInterval != 0 && i%c.StoreInterval == 0 {
			backup.Save()
		}
		time.Sleep(time.Second)
		i++
		if i >= math.MaxInt32 {
			i = 0
		}
	}
}
