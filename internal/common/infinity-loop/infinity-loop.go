package infinity_loop

import "time"

func Run() {
	for {
		select {
		case <-time.After(time.Second):
		}
	}
}
