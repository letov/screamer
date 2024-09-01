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
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		ss.backupService.Save()
		os.Exit(1)
	}()
}

func NewShutdownService(bs *BackupService) *ShutdownService {
	return &ShutdownService{
		backupService: bs,
	}
}
