package app

import (
	event_loop "screamer/internal/common/infrastructure/eventloop"
	"screamer/internal/common/infrastructure/prof"
)

func Start(*event_loop.EventLoop, *prof.Server) {}
