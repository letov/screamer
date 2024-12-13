package app

import (
	"net/http"
	event_loop "screamer/internal/common/eventloop"
)

func Start(*event_loop.EventLoop, *http.Server) {}
