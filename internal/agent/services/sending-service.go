package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"screamer/internal/agent/config"
	"screamer/internal/agent/repositories"
	"screamer/internal/common/metric"
)

type SendingService struct {
	config *config.Config
	repo   repositories.Repository
	log    *zap.SugaredLogger
}

func (ss *SendingService) SendMetrics(ctx context.Context) {
	ms := ss.repo.GetAll(ctx)

	for _, m := range ms {
		ss.requestOne(m)
	}
}

func (ss *SendingService) requestOne(m metric.Metric) {
	url := fmt.Sprintf("http://%v/update", ss.config.NetAddress.String())
	body, _ := m.Bytes()

	r, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err == nil {
		defer func(Body io.ReadCloser) {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = Body.Close()
		}(r.Body)
	}

	if err != nil {
		ss.log.Warn("Request error", err.Error())
	} else if r.StatusCode != http.StatusOK {
		ss.log.Warn("Bad status", r.StatusCode)
	} else {
		data, _ := io.ReadAll(r.Body)
		ss.log.Info("Answer: ", string(data))
	}
}

func (ss *SendingService) requestAll(ms []metric.Metric) {
	url := fmt.Sprintf("http://%v/updates", ss.config.NetAddress.String())
	var jms []metric.JSONMetric

	for _, m := range ms {
		jm, err := m.JSON()
		if err != nil {
			ss.log.Warn("Request error", err.Error())
			return
		}
		jms = append(jms, jm)
	}

	body, err := json.Marshal(jms)
	if err != nil {
		ss.log.Warn("Request error", err.Error())
		return
	}

	r, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err == nil {
		defer func(Body io.ReadCloser) {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = Body.Close()
		}(r.Body)
	}

	if err != nil {
		ss.log.Warn("Request error", err.Error())
	} else if r.StatusCode != http.StatusOK {
		ss.log.Warn("Bad status", r.StatusCode)
	} else {
		data, _ := io.ReadAll(r.Body)
		ss.log.Info("Answer: ", string(data))
	}
}

func NewSendingService(log *zap.SugaredLogger, config *config.Config, repo repositories.Repository) *SendingService {
	return &SendingService{
		config: config,
		repo:   repo,
		log:    log,
	}
}
