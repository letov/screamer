package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"screamer/internal/agent/config"
	"screamer/internal/agent/repositories"
	"screamer/internal/common/metric"
)

type SendingService struct {
	config *config.Config
	repo   repositories.Repository
}

func (ss *SendingService) SendMetrics() {
	url := fmt.Sprintf("%v/update", ss.config.NetAddress.String())
	ms := ss.repo.GetAll()

	for _, m := range ms {
		ss.request(url, m)
	}
}

func (ss *SendingService) request(url string, m metric.Metric) {
	j, _ := m.Json()
	body, _ := json.Marshal(&j)

	r, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err == nil {
		defer func(Body io.ReadCloser) {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = Body.Close()
		}(r.Body)
	}

	if ss.config.AgentLogEnable {
		if err != nil {
			log.Println("Request error", err.Error())
		} else if r.StatusCode != http.StatusOK {
			log.Println("Bad status", r.StatusCode)
		} else {
			data, _ := io.ReadAll(r.Body)
			log.Println("Answer", string(data))
		}
	}
}

func NewSendingService(config *config.Config, repo repositories.Repository) *SendingService {
	return &SendingService{
		config: config,
		repo:   repo,
	}
}
