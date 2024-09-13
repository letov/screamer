package repositories

import (
	"context"
	"encoding/json"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
	"screamer/internal/common/metric"
	"screamer/internal/server/config"
	"sync"
)

type FileRepository struct {
	c  *config.Config
	l  *zap.SugaredLogger
	mr *MemoryRepository
	sync.Mutex
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
		err = fr.toFile(fr.GetAll(ctx))
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
		err = fr.toFile(fr.GetAll(ctx))
	}
	return newM, err
}

func (fr *FileRepository) SaveAllToFile(ctx context.Context) {
	err := fr.toFile(fr.GetAll(ctx))
	fr.processError(err)
}

func (fr *FileRepository) LoadAllFromFile(ctx context.Context) {
	ms, err := fr.fromFile()
	if err != nil {
		fr.l.Warn("Load form file error:", err.Error())
		return
	}
	for _, m := range ms {
		_, err := fr.mr.Add(ctx, *m)
		fr.processError(err)
	}
}

func (fr *FileRepository) toFile(ms []metric.Metric) error {
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
	return err
}

func (fr *FileRepository) fromFile() ([]*metric.Metric, error) {
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
