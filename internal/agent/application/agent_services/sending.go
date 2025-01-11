package agent_services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"screamer/internal/agent/application/repo"
	"screamer/internal/agent/infrastructure/config"
	"screamer/internal/common"
	"screamer/internal/common/application/dto"
	"screamer/internal/common/domain"
	"screamer/internal/common/helpers/hash"
	"screamer/internal/common/helpers/retry"
	"screamer/internal/common/infrastructure/grpcclient"
	"sync/atomic"
	"time"

	"github.com/aoliveti/curling"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Sending struct {
	config     *config.Config
	repo       repo.Repository
	log        *zap.SugaredLogger
	encrypt    *hash.RSAEncrypt
	activeJobs atomic.Int32
	gc         *grpcclient.GRPCClient
}

func (ss *Sending) SendMetrics(ctx context.Context) {
	ms := ss.repo.GetAll(ctx)

	jobs := make(chan domain.Metric, len(ms))

	for w := 0; w < max(ss.config.RateLimit, 1); w++ {
		go ss.worker(ctx, jobs)
	}

	for _, m := range ms {
		jobs <- m
	}

	close(jobs)
}

func (ss *Sending) worker(ctx context.Context, jobs <-chan domain.Metric) {
	for j := range jobs {
		ss.requestOne(ctx, j)
	}
}

func (ss *Sending) requestOne(ctx context.Context, m domain.Metric) {
	url := fmt.Sprintf("http://%v/update", ss.config.NetAddress.String())
	body, _ := json.Marshal(m)
	if ss.encrypt != nil {
		body, _ = ss.encrypt.Encrypt(body)
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		cancel()
	}()

	var job func(ctx context.Context) ([]byte, error)
	if len(ss.config.NetAddressGrpc.String()) > 0 {
		job = ss.requestJobGrpc(m, ss.gc, &ss.activeJobs)
	} else {
		job = ss.requestJobHTTP(&body, url, &ss.activeJobs)
	}

	_, _ = retry.NewRetryJob(ctxWithTimeout, "agent request", job, []error{}, []int{1, 2, 5}, ss.log)
}

func (ss *Sending) getIP() string {
	ips, _ := net.LookupIP(ss.config.Host)
	if len(ips) == 0 {
		ss.log.Fatal("host lookup fail")
	}
	return ips[0].String()
}

func (ss *Sending) requestAll(ctx context.Context, ms []domain.Metric) {
	url := fmt.Sprintf("http://%v/updates", ss.config.NetAddress.String())
	var jms []dto.JSONMetric

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for _, m := range ms {
		jm, err := dto.NewJSONMetric(m)
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

	job := ss.requestJobHTTP(&body, url, &ss.activeJobs)
	_, _ = retry.NewRetryJob(ctxWithTimeout, "agent request", job, []error{}, []int{1, 2, 5}, ss.log)
}

func (ss *Sending) requestJobHTTP(
	body *[]byte,
	url string,
	aj *atomic.Int32,
) func(ctx context.Context) ([]byte, error) {
	return func(ctx context.Context) ([]byte, error) {
		aj.Add(1)
		defer func() {
			aj.Add(-1)
		}()

		client := http.Client{}
		ip := ss.getIP()
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(*body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Real-IP", ip)
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

func (ss *Sending) requestJobGrpc(
	m domain.Metric,
	gc *grpcclient.GRPCClient,
	aj *atomic.Int32,
) func(ctx context.Context) ([]byte, error) {
	return func(ctx context.Context) ([]byte, error) {
		aj.Add(1)
		defer func() {
			aj.Add(-1)
		}()

		in, err := dto.NewPbMetric(m)
		if err != nil {
			return nil, err
		}

		_, err = gc.Client.UpdateValue(ctx, in)

		return nil, err
	}
}

func NewSending(
	lc fx.Lifecycle,
	log *zap.SugaredLogger,
	config *config.Config,
	repo repo.Repository,
	gc *grpcclient.GRPCClient,
) *Sending {
	var encrypt *hash.RSAEncrypt
	if len(config.CryptoKey) != 0 {
		encrypt = hash.NewRSAEncrypt(config.CryptoKey, log)
	}

	ss := &Sending{
		config:  config,
		repo:    repo,
		log:     log,
		encrypt: encrypt,
		gc:      gc,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			aj := ss.activeJobs.Load()
			log.Info("Agent active jobs count: ", aj)
			if aj != 0 {
				for i := 0; i < 5; i++ {
					aj = ss.activeJobs.Load()
					if aj > 0 {
						log.Info("Agent active jobs count: ", aj)
						log.Info("Try to wait")
						time.Sleep(time.Second)
					} else {
						break
					}
				}
			}
			log.Info("Sending service closed")
			return nil
		},
	})

	return ss
}
