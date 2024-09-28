package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aoliveti/curling"
	"go.uber.org/zap"
	"io"
	"net/http"
	"screamer/internal/agent/config"
	"screamer/internal/agent/repositories"
	"screamer/internal/common"
	"screamer/internal/common/hash"
	"screamer/internal/common/metric"
	"screamer/internal/common/retry"
	"time"
)

type SendingService struct {
	config *config.Config
	repo   repositories.Repository
	log    *zap.SugaredLogger
}

func (ss *SendingService) SendMetrics(ctx context.Context) {
	ms := ss.repo.GetAll(ctx)

	for _, m := range ms {
		ss.requestOne(ctx, m)
	}
}

func (ss *SendingService) requestOne(ctx context.Context, m metric.Metric) {
	url := fmt.Sprintf("http://%v/update", ss.config.NetAddress.String())
	body, _ := m.Bytes()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	job := ss.requestJob(&body, url)
	_, _ = retry.NewRetryJob(ctxWithTimeout, "agent request", job, []error{}, []int{1, 2, 5}, ss.log)
}

func (ss *SendingService) requestAll(ctx context.Context, ms []metric.Metric) {
	url := fmt.Sprintf("http://%v/updates", ss.config.NetAddress.String())
	var jms []metric.JSONMetric

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

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

	job := ss.requestJob(&body, url)
	_, _ = retry.NewRetryJob(ctxWithTimeout, "agent request", job, []error{}, []int{1, 2, 5}, ss.log)
}

func (ss *SendingService) requestJob(body *[]byte, url string) func(ctx context.Context) ([]byte, error) {
	return func(ctx context.Context) ([]byte, error) {
		client := http.Client{}
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(*body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("HashSHA256", hash.Encode(body, ss.config.Key))
		cmd, _ := curling.NewFromRequest(req)
		ss.log.Info(cmd)
		res, err := client.Do(req)
		if err == nil {
			defer func(Body io.ReadCloser) {
				_, _ = io.Copy(io.Discard, res.Body)
				_ = Body.Close()
			}(res.Body)
		}

		if err != nil {
			return nil, err
		}
		resBody, err := io.ReadAll(res.Body)
		_ = res.Body.Close()
		if res.StatusCode != http.StatusOK {
			err = common.ErrNoOKResponse
		}
		return resBody, err
	}
}

func NewSendingService(log *zap.SugaredLogger, config *config.Config, repo repositories.Repository) *SendingService {
	return &SendingService{
		config: config,
		repo:   repo,
		log:    log,
	}
}
