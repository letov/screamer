package store

import (
	"context"
	"encoding/json"
	"os"
	"screamer/internal/common/application/dto"
	"screamer/internal/common/domain"
	"screamer/internal/common/helpers/retry"
	"screamer/internal/server/infrastructure/config"
	"sync"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type File struct {
	c  *config.Config
	l  *zap.SugaredLogger
	mr *Memory
	sync.Mutex
}

func (fr *File) BatchUpdate(_ context.Context, _ []domain.Metric) error {
	//TODO implement me
	panic("implement me")
}

type JSONMetricList struct {
	Array []dto.JSONMetric
}

func (fr *File) GetAll(ctx context.Context) []domain.Metric {
	return fr.mr.GetAll(ctx)
}

func (fr *File) Add(ctx context.Context, m domain.Metric) (domain.Metric, error) {
	newM, err := fr.mr.Add(ctx, m)
	if err != nil {
		return newM, err
	}
	if fr.c.Restore && fr.c.StoreInterval == 0 {
		err = fr.toFile(ctx, fr.GetAll(ctx))
	}
	return newM, err
}

func (fr *File) Get(ctx context.Context, i domain.Ident) (domain.Metric, error) {
	return fr.mr.Get(ctx, i)
}

func (fr *File) Increase(ctx context.Context, m domain.Metric) (domain.Metric, error) {
	newM, err := fr.mr.Increase(ctx, m)
	if err != nil {
		return newM, err
	}
	if fr.c.Restore && fr.c.StoreInterval == 0 {
		err = fr.toFile(ctx, fr.GetAll(ctx))
	}
	return newM, err
}

func (fr *File) SaveAllToFile(ctx context.Context) {
	err := fr.toFile(ctx, fr.GetAll(ctx))
	fr.processError(err)
}

func (fr *File) LoadAllFromFile(ctx context.Context) {
	ms, err := fr.fromFile(ctx)
	if err != nil {
		fr.l.Warn("Load form file error:", err.Error())
		return
	}
	for _, m := range ms {
		_, err := fr.mr.Add(ctx, m)
		fr.processError(err)
	}
}

func (fr *File) toFile(ctx context.Context, ms []domain.Metric) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	job := fr.toFileJob(ms)
	_, err := retry.NewRetryJob(ctxWithTimeout, "wait file access", job, []error{}, []int{1, 2, 5}, fr.l)
	return err
}

func (fr *File) toFileJob(ms []domain.Metric) func(ctx context.Context) (bool, error) {
	return func(ctx context.Context) (bool, error) {
		fp := fr.c.FileStoragePath

		jms := make([]dto.JSONMetric, 0)
		for _, m := range ms {
			j, err := dto.NewJSONMetric(m)
			fr.processError(err)
			jms = append(jms, j)
		}

		jml := &JSONMetricList{Array: jms}
		body, err := json.MarshalIndent(jml, "", "   ")
		fr.processError(err)

		fr.Lock()
		err = os.WriteFile(fp, body, 0777)
		fr.Unlock()
		return true, err
	}
}

func (fr *File) fromFile(ctx context.Context) ([]domain.Metric, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	job := fr.fromFileJob()
	return retry.NewRetryJob(ctxWithTimeout, "wait file access", job, []error{}, []int{1, 2, 5}, fr.l)
}

func (fr *File) fromFileJob() func(ctx context.Context) ([]domain.Metric, error) {
	return func(ctx context.Context) ([]domain.Metric, error) {
		fp := fr.c.FileStoragePath

		data, err := os.ReadFile(fp)
		if err != nil {
			return nil, err
		}

		jml := &JSONMetricList{}
		err = json.Unmarshal(data, jml)
		if err != nil {
			return nil, err
		}

		res := make([]domain.Metric, 0)
		for _, jm := range jml.Array {
			m, err := jm.GetDomainMetric()
			fr.processError(err)
			res = append(res, m)
		}

		return res, nil
	}
}

func (fr *File) processError(err error) {
	if err != nil {
		fr.l.Warn("File process error:", err.Error())
	}
}

func NewFile(lc fx.Lifecycle, c *config.Config, log *zap.SugaredLogger, mr *Memory) *File {
	fr := &File{
		c:  c,
		l:  log,
		mr: mr,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if c.Restore {
				log.Info("Loading from file")
				fr.LoadAllFromFile(ctx)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if c.Restore {
				log.Info("Saving to file")
				fr.SaveAllToFile(ctx)
			}
			return nil
		},
	})

	return fr
}
