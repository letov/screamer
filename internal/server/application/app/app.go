package app

import (
	"net/http"
	event_loop "screamer/internal/common/infrastructure/eventloop"

	"google.golang.org/grpc"
)

func Start(*event_loop.EventLoop, *http.Server, *grpc.Server) {}
