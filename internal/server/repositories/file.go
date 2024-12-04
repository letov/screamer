package repositories

import (
	"context"
	"encoding/json"
	"os"
	"screamer/internal/common/metric"
	"screamer/internal/common/retry"
	"screamer/internal/server/config"
	"sync"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type FileRepository struct {
	c  *config.Config
	l  *zap.SugaredLogger
	mr *MemoryRepository
	sync.Mutex
}

func (fr *FileRepository) BatchUpdate(_ context.Context, _ []metric.Metric) error {
	//TODO implement me
	panic("implement me")
}

type JSONMetricList struct {
	Array []metric.JSONMetric
}

func (fr *FileRepository) GetAll(ctx context.Context) []metric.Metric {
	return fr.mr.GetAll(ctx)
}

func (fr *FileRepository) Add(ctx context.Context, m metric.Metric) (metric.Metric, error) {
	newM, err := fr.mr.Add(ctx, m)
	if err != nil {
		return newM, err
	}
	if fr.c.Restore && fr.c.StoreInterval == 0 {
		err = fr.toFile(ctx, fr.GetAll(ctx))
	}
	return newM, err
}

func (fr *FileRepository) Get(ctx context.Context, i metric.Ident) (metric.Metric, error) {
	return fr.mr.Get(ctx, i)
}

func (fr *FileRepository) Increase(ctx context.Context, m metric.Metric) (metric.Metric, error) {
	newM, err := fr.mr.Increase(ctx, m)
	if err != nil {
		return newM, err
	}
	if fr.c.Restore && fr.c.StoreInterval == 0 {
		err = fr.toFile(ctx, fr.GetAll(ctx))
	}
	return newM, err
}

func (fr *FileRepository) SaveAllToFile(ctx context.Context) {
	err := fr.toFile(ctx, fr.GetAll(ctx))
	fr.processError(err)
}

func (fr *FileRepository) LoadAllFromFile(ctx context.Context) {
	ms, err := fr.fromFile(ctx)
	if err != nil {
		fr.l.Warn("Load form file error:", err.Error())
		return
	}
	for _, m := range ms {
		_, err := fr.mr.Add(ctx, *m)
		fr.processError(err)
	}
}

func (fr *FileRepository) toFile(ctx context.Context, ms []metric.Metric) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	job := fr.toFileJob(ms)
	_, err := retry.NewRetryJob(ctxWithTimeout, "wait file access", job, []error{}, []int{1, 2, 5}, fr.l)
	return err
}

func (fr *FileRepository) toFileJob(ms []metric.Metric) func(ctx context.Context) (bool, error) {
	return func(ctx context.Context) (bool, error) {
		fp := fr.c.FileStoragePath

		jms := make([]metric.JSONMetric, 0)
		for _, m := range ms {
			j, err := m.JSON()
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

func (fr *FileRepository) fromFile(ctx context.Context) ([]*metric.Metric, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	job := fr.fromFileJob()
	return retry.NewRetryJob(ctxWithTimeout, "wait file access", job, []error{}, []int{1, 2, 5}, fr.l)
}

func (fr *FileRepository) fromFileJob() func(ctx context.Context) ([]*metric.Metric, error) {
	return func(ctx context.Context) ([]*metric.Metric, error) {
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

		res := make([]*metric.Metric, 0)
		for _, jm := range jml.Array {
			m, err := metric.NewMetricFromJSON(&jm)
			fr.processError(err)
			res = append(res, m)
		}

		return res, nil
	}
}

func (fr *FileRepository) processError(err error) {
	if err != nil {
		fr.l.Warn("File process error:", err.Error())
	}
}

func NewFileRepository(lc fx.Lifecycle, c *config.Config, log *zap.SugaredLogger, mr *MemoryRepository) *FileRepository {
	fr := &FileRepository{
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
