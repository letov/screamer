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
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)
	done := make(chan bool, 1)

	go func() {
		<-sigs
		ss.backupService.Save()
		done <- true
	}()

	<-done
}

func NewShutdownService(bs *BackupService) *ShutdownService {
	return &ShutdownService{
		backupService: bs,
	}
}
