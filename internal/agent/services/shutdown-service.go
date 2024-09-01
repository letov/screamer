package services

import (
	"os"
	"os/signal"
	"syscall"
)

type ShutdownService struct {
}

func (ss *ShutdownService) Run() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		<-sigs

		done <- true
	}()

	<-done
}

func NewShutdownService() *ShutdownService {
	return &ShutdownService{}
}
