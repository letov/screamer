package services

type LoadService struct {
	backupService *BackupService
}

func (ls *LoadService) Run() {
	ls.backupService.Load()
}

func NewLoadService(bs *BackupService) *LoadService {
	return &LoadService{
		backupService: bs,
	}
}
