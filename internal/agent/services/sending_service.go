package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"screamer/internal/agent/config"
	"screamer/internal/agent/repositories"
	"screamer/internal/common"
	"screamer/internal/common/hash"
	"screamer/internal/common/metric"
	"screamer/internal/common/retry"
	"sync"
	"time"

	"github.com/aoliveti/curling"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type SendingService struct {
	config  *config.Config
	repo    repositories.Repository
	log     *zap.SugaredLogger
	encrypt *hash.RSAEncrypt
	stop    bool
	sync.Mutex
	activeJobs int
}

func (ss *SendingService) SendMetrics(ctx context.Context) {
	if ss.stop {
		return
	}
	ms := ss.repo.GetAll(ctx)

	jobs := make(chan metric.Metric, len(ms))

	for w := 0; w < max(ss.config.RateLimit, 1); w++ {
		go ss.worker(ctx, jobs)
	}

	for _, m := range ms {
		jobs <- m
	}

	close(jobs)
}

func (ss *SendingService) worker(ctx context.Context, jobs <-chan metric.Metric) {
	for j := range jobs {
		ss.requestOne(ctx, j)
	}
}

func (ss *SendingService) requestOne(ctx context.Context, m metric.Metric) {
	ss.Lock()
	ss.activeJobs++
	ss.Unlock()
	url := fmt.Sprintf("http://%v/update", ss.config.NetAddress.String())
	body, _ := m.Bytes()
	if ss.encrypt != nil {
		body, _ = ss.encrypt.Encrypt(body)
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		ss.Lock()
		ss.activeJobs--
		ss.Unlock()
		cancel()
	}()

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

func NewSendingService(
	lc fx.Lifecycle,
	log *zap.SugaredLogger,
	config *config.Config,
	repo repositories.Repository,
) *SendingService {
	var encrypt *hash.RSAEncrypt
	if len(config.CryptoKey) != 0 {
		encrypt = hash.NewRSAEncrypt(config.CryptoKey, log)
	}

	ss := &SendingService{
		config:     config,
		repo:       repo,
		log:        log,
		encrypt:    encrypt,
		activeJobs: 0,
		stop:       false,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			ss.stop = true
			log.Info("Agent active jobs count: ", ss.activeJobs)
			for i := 0; i < 5; i++ {
				if ss.activeJobs > 0 {
					log.Info("Agent active jobs count: ", ss.activeJobs)
					log.Info("Try to wait sending")
					time.Sleep(time.Second)
				} else {
					log.Info("Try to wait")
					break
				}
			}
			return nil
		},
	})

	return ss
}
