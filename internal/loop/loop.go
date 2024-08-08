package loop

import (
	"math"
	"screamer/internal/collector"
	"screamer/internal/config"
	"screamer/internal/pusher"
	"time"
)

func Run() {
	c := config.GetConfig()
	var i = 0
	for {
		if i%c.PollInterval == 0 {
			collector.UpdateMetrics()
		}
		if i%c.ReportInterval == 0 {
			pusher.SendData()
		}
		time.Sleep(time.Second)
		i++
		if i >= math.MaxInt32 {
			i = 0
		}
	}
}
