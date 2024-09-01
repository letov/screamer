package services

import (
	"os"
	"os/signal"
	"syscall"
)

type ShutdownService struct {
	backupService *BackupService
}

func (ss *ShutdownService) Run() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGHUP:
				ss.backupService.Save()

			case syscall.SIGINT:
				ss.backupService.Save()

			case syscall.SIGTERM:
				ss.backupService.Save()
				exit_chan <- 0

			case syscall.SIGQUIT:
				ss.backupService.Save()
				exit_chan <- 0
			}
		}
	}()

	code := <-exit_chan
	os.Exit(code)
}

func NewShutdownService(bs *BackupService) *ShutdownService {
	return &ShutdownService{
		backupService: bs,
	}
}
